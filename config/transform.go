package config

import (
	"errors"
	"os"
)

// Creating error vars like this
// We'll be able to use it to check what kind of error without
// reading the content of error.
// e.g: if (err == config.AVRO_SCHEMA_URL)
var (
	invalidUrl = errors.New("no Avro url AvroConfig.Url defined, please set the AVRO_SCHEMA_URL env")
)

func GetAvro() (config EnvAvroConfig, err error) {
	avroConfig := EnvAvroConfig{
		Url:      os.Getenv("AVRO_SCHEMA_URL"),
		Username: os.Getenv("AVRO_SCHEMA_USERNAME"),
		Password: os.Getenv("AVRO_SCHEMA_PASSWORD"),
	}

	if len(avroConfig.Url) == 0 {
		return avroConfig, invalidUrl
	}

	return avroConfig, nil
}

// Creating error vars like this
// We'll be able to use it to check what kind of error without
// reading the content of error.
// e.g: if (err == config.invalidBroker)
var (
	invalidBroker        = errors.New("no Kafka bootstrap kafkaConfig.Brokers defined, please set the KAFKA_BROKERS env")
	invalidTopic         = errors.New("no kafkaConfig.Topic given to be consumed, please set the KAFKA_TOPICS env")
	invalidConsumerGroup = errors.New("no Kafka consumer kafkaConfig.Group defined, please set the KAFKA_GROUP env")
	invalidAuthSsl       = errors.New("not enough SSL Auth config defined, please set KAFKA_AUTHENTICATION_CA, KAFKA_AUTHENTICATION_CERTIFICATE, KAFKA_AUTHENTICATION_KEY envs")
	invalidAuthSaslSsl   = errors.New("not enough sasl_ssl Auth config defined, please set KAFKA_USERNAME, KAFKA_PASSWORD envs")
)

func GetKafka() (config EnvKafkaConfig, err error) {
	kafkaConfig := EnvKafkaConfig{
		Brokers:         os.Getenv("KAFKA_BROKERS"),
		Version:         os.Getenv("KAFKA_VERSION"),
		ConsumerGroup:   os.Getenv("KAFKA_GROUP"),
		Topic:           os.Getenv("KAFKA_TOPICS"),
		Assignor:        os.Getenv("KAFKA_ASSIGNOR"),
		OldestFirst:     true,
		AuthType:        os.Getenv("KAFKA_AUTHENTICATION_TYPE"),
		AuthCa:          os.Getenv("KAFKA_AUTHENTICATION_CA"),
		AuthKey:         os.Getenv("KAFKA_AUTHENTICATION_KEY"),
		AuthCertificate: os.Getenv("KAFKA_AUTHENTICATION_CERTIFICATE"),
		Username:        os.Getenv("KAFKA_USERNAME"),
		Password:        os.Getenv("KAFKA_PASSWORD"),
	}

	if len(kafkaConfig.Brokers) == 0 {
		return kafkaConfig, invalidBroker
	}

	if len(kafkaConfig.Topic) == 0 {
		return kafkaConfig, invalidTopic
	}

	if len(kafkaConfig.ConsumerGroup) == 0 {
		return kafkaConfig, invalidConsumerGroup
	}

	if kafkaConfig.AuthType == "ssl" {
		if len(kafkaConfig.AuthKey) == 0 {
			return kafkaConfig, invalidAuthSsl
		}

		if len(kafkaConfig.AuthCa) == 0 {
			return kafkaConfig, invalidAuthSsl
		}

		if len(kafkaConfig.AuthCertificate) == 0 {
			return kafkaConfig, invalidAuthSsl
		}
	}

	if kafkaConfig.AuthType == "sasl_ssl" {
		if len(kafkaConfig.Username) == 0 {
			return kafkaConfig, invalidAuthSaslSsl
		}

		if len(kafkaConfig.Password) == 0 {
			return kafkaConfig, invalidAuthSaslSsl
		}
	}

	return kafkaConfig, nil
}
