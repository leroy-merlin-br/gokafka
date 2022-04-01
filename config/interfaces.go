package config

import "github.com/Shopify/sarama"

type Config struct {
	Kafka         *sarama.Config
	Brokers       string
	ConsumerGroup string
	Topic         string
}

type EnvKafkaConfig struct {
	Brokers         string
	Version         string
	ConsumerGroup   string
	Topic           string
	Assignor        string
	OldestFirst     bool
	AuthType        string
	AuthCa          string
	AuthCertificate string
	AuthKey         string
	Username        string
	Password        string
}

type EnvAvroConfig struct {
	Url      string
	Username string
	Password string
}
