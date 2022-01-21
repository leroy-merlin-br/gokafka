package kafka

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"lmbr/merchant/config"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func Handle(consumer Consumer) (err error) {
	kafkaConfig, err := config.GetKafka()
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

func consume(client sarama.ConsumerGroup, wg *sync.WaitGroup, kafkaConfig config.KafkaConfig, consumer Consumer, ctx context.Context) <-chan error {
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
