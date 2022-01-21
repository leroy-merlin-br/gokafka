package kafka

import (
	"crypto/tls"
	"github.com/Shopify/sarama"
	"github.com/leroy-merlin-br/lmbr-kafka/config"
	"github.com/pkg/errors"
)

func NewConfig(kafkaConfig config.KafkaConfig) (*sarama.Config, error) {
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

func authentication(saramaConfig *sarama.Config, kafkaConfig config.KafkaConfig) error {
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

func saslSslAuthentication(saramaConfig *sarama.Config, kafkaConfig config.KafkaConfig) {
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

func sslAuthentication(saramaConfig *sarama.Config, kafkaConfig config.KafkaConfig) error {
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
