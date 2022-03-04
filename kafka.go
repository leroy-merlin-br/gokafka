package gokafka

import (
	"bytes"
	"context"
	"github.com/Shopify/sarama"
	"github.com/leroy-merlin-br/gokafka/config"
	"github.com/linkedin/goavro"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"sync"
	"syscall"
)

func Handle(consumer Consumer) (err error) {
	config, err := config.Make()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := sarama.NewConsumerGroup(strings.Split(config.Brokers, ","), config.ConsumerGroup, config.Kafka)
	defer func() {
		err = client.Close()
	}()

	if err != nil {
		return errors.Wrap(err, "Error creating consumer group client: %v")
	}
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(1)

	occurancesErr := consume(client, wg, *config, consumer, ctx)

	<-consumer.Ready // Await till the consumer has been set up
	log.Print("Consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Print("terminating: context cancelled")
	case <-sigterm:
		log.Print("terminating: via signal")
	}

	var open bool
	if err, open = <-occurancesErr; open {
		return
	}

	return err
}

func consume(client sarama.ConsumerGroup, wg *sync.WaitGroup, kafkaConfig config.Config, consumer Consumer, ctx context.Context) <-chan error {
	errs := make(chan error, 1)
	defer close(errs)

	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, strings.Split(kafkaConfig.Topic, ","), &consumer); err != nil {
				errs <- errors.Wrap(err, "Error creating consumer group client: %v")
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.Ready = make(chan bool)
		}
	}()

	return errs
}

type Action func(record *goavro.Record) error

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	Ready  chan bool
	Action Action
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.Ready)

	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	codec, err := Codec()
	if err != nil {
		return err
	}

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		bb := bytes.NewBuffer(message.Value[5:])
		decoded, _ := codec.Decode(bb)
		record, ok := decoded.(*goavro.Record)
		if !ok {
			return errors.New("Type: " + reflect.TypeOf(decoded).String() + "is not a valid Record.")
		}

		err := consumer.Action(record)
		if err != nil {
			return err
		}

		session.MarkMessage(message, "")
	}

	return nil
}
