package gokafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/leroy-merlin-br/gokafka/config"
	"github.com/leroy-merlin-br/gokafka/consumer"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

func Handle(consumer consumer.Consumer) (err error) {
	consumerConfig, err := config.Make()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := sarama.NewConsumerGroup(strings.Split(consumerConfig.Brokers, ","), consumerConfig.ConsumerGroup, consumerConfig.Kafka)
	defer func() {
		err = client.Close()
	}()

	if err != nil {
		return errors.Wrap(err, "Error creating consumer group client: %v")
	}
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(1)

	occurancesErr := consume(client, wg, *consumerConfig, consumer, ctx)

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

func consume(client sarama.ConsumerGroup, wg *sync.WaitGroup, consumerConfig config.Config, consumer consumer.Consumer, ctx context.Context) <-chan error {
	errs := make(chan error, 1)
	defer close(errs)

	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, strings.Split(consumerConfig.Topic, ","), &consumer); err != nil {
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
