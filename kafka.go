package gokafka

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/Shopify/sarama"
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

func NewConfig(kafkaConfig KafkaConfig) (*sarama.Config, error) {
	saramaConfig := sarama.NewConfig()

	err := authentication(saramaConfig, kafkaConfig)
	if err != nil {
		return saramaConfig, errors.Wrap(err, "Error parsing Kafka authentication: %v")
	}

	version, err := sarama.ParseKafkaVersion(kafkaConfig.Version)
	if err != nil {
		return saramaConfig, errors.Wrap(err, "Error parsing Kafka version: %v")
	}

	saramaConfig.Version = version

	switch kafkaConfig.Assignor {
	case "sticky":
		saramaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		saramaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		saramaConfig.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		err = errors.New("Unrecognized consumer group partition assignor: %s")
	}

	if kafkaConfig.OldestFirst {
		saramaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	return saramaConfig, err
}

func authentication(saramaConfig *sarama.Config, kafkaConfig KafkaConfig) error {
	if kafkaConfig.AuthType == "ssl" {
		err := sslAuthentication(saramaConfig, kafkaConfig)
		if err != nil {
			return err
		}

		return nil
	}

	if kafkaConfig.AuthType == "sasl_ssl" {
		saslSslAuthentication(saramaConfig, kafkaConfig)
	}

	return nil
}

func saslSslAuthentication(saramaConfig *sarama.Config, kafkaConfig KafkaConfig) {
	saramaConfig.Net.SASL.User = kafkaConfig.Username
	saramaConfig.Net.SASL.Password = kafkaConfig.Password
	saramaConfig.Net.SASL.Handshake = true
	saramaConfig.Net.SASL.Enable = true
	saramaConfig.Net.SASL.Mechanism = sarama.SASLTypePlaintext

	saramaConfig.Net.TLS.Enable = true
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ClientAuth:         0,
	}
	tlsConfig.InsecureSkipVerify = true

	saramaConfig.Net.TLS.Config = tlsConfig
}

func sslAuthentication(saramaConfig *sarama.Config, kafkaConfig KafkaConfig) error {
	tlsConfig, err := NewTLSConfig(
		kafkaConfig.AuthCertificate,
		kafkaConfig.AuthKey,
		kafkaConfig.AuthCa)
	if err != nil {
		return err
	}

	// This can be used on test server if domain does not match cert:
	tlsConfig.InsecureSkipVerify = true

	saramaConfig.Net.TLS.Enable = true
	saramaConfig.Net.TLS.Config = tlsConfig

	return nil
}

func Handle(consumer Consumer) (err error) {
	kafkaConfig, err := GetKafka()
	if err != nil {
		return err
	}

	saramaConfig, err := NewConfig(kafkaConfig)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := sarama.NewConsumerGroup(strings.Split(kafkaConfig.Brokers, ","), kafkaConfig.ConsumerGroup, saramaConfig)
	defer func() {
		err = client.Close()
	}()

	if err != nil {
		return errors.Wrap(err, "Error creating consumer group client: %v")
	}
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	wg.Add(1)

	occurancesErr := consume(client, wg, kafkaConfig, consumer, ctx)

	<-consumer.Ready // Await till the consumer has been set up
	log.Print("Sarama consumer up and running!...")

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

func consume(client sarama.ConsumerGroup, wg *sync.WaitGroup, kafkaConfig KafkaConfig, consumer Consumer, ctx context.Context) <-chan error {
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

// NewTLSConfig generates a TLS configuration used to authenticate on server with
// certificates.
// Parameters are the three pem files path we need to authenticate: client cert, client key and CA cert.
func NewTLSConfig(cert string, key string, ca string) (*tls.Config, error) {
	tlsConfig := tls.Config{}

	// Load client cert
	certificate, err := tls.X509KeyPair([]byte(cert), []byte(key))
	if err != nil {
		return &tlsConfig, err
	}
	tlsConfig.Certificates = []tls.Certificate{certificate}

	// Load CA cert
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM([]byte(ca))
	tlsConfig.RootCAs = caCertPool

	tlsConfig.BuildNameToCertificate()
	return &tlsConfig, err
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
