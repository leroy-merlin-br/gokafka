package config

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

func Make() (*Config, error) {
	envKafkaConfig, err := GetKafka()
	if err != nil {
		return nil, err
	}

	saramaConfig := sarama.NewConfig()

	err = ConfigureSarama(envKafkaConfig, saramaConfig)
	if err != nil {
		return nil, err
	}

	config := &Config{
		Kafka:         saramaConfig,
		Brokers:       envKafkaConfig.Brokers,
		ConsumerGroup: envKafkaConfig.ConsumerGroup,
		Topic:         envKafkaConfig.Topic,
	}

	return config, nil
}

func ConfigureSarama(kafkaConfig EnvKafkaConfig, saramaConfig *sarama.Config) error {
	err := authentication(saramaConfig, kafkaConfig)
	if err != nil {
		return errors.Wrap(err, "Error parsing Kafka authentication: %v")
	}

	version, err := sarama.ParseKafkaVersion(kafkaConfig.Version)
	if err != nil {
		return errors.Wrap(err, "Error parsing Kafka version: %v")
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

	return nil
}

func authentication(saramaConfig *sarama.Config, kafkaConfig EnvKafkaConfig) error {
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

func saslSslAuthentication(saramaConfig *sarama.Config, kafkaConfig EnvKafkaConfig) {
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

func sslAuthentication(saramaConfig *sarama.Config, kafkaConfig EnvKafkaConfig) error {
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
