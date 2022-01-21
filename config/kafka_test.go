package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetKafkaConfig(t *testing.T) {
	os.Setenv("KAFKA_BROKERS", "test_broker")
	os.Setenv("KAFKA_VERSION", "2.1.1")
	os.Setenv("KAFKA_GROUP", "test_unit")
	os.Setenv("KAFKA_TOPICS", "test_topic")
	os.Setenv("KAFKA_ASSIGNOR", "range")
	os.Setenv("KAFKA_AUTHENTICATION_TYPE", "none")

	kafkaConfig, err := GetKafka()

	assert.Empty(t, err)
	assert.Equal(t, "range", kafkaConfig.Assignor)
	assert.Equal(t, "test_topic", kafkaConfig.Topic)
	assert.Equal(t, "test_unit", kafkaConfig.ConsumerGroup)
	assert.Equal(t, "2.1.1", kafkaConfig.Version)
	assert.Equal(t, "test_broker", kafkaConfig.Brokers)
	assert.Equal(t, "none", kafkaConfig.AuthType)
}

func TestGetKafkaConfigWithSSLAuth(t *testing.T) {
	os.Setenv("KAFKA_BROKERS", "test_broker")
	os.Setenv("KAFKA_VERSION", "2.1.1")
	os.Setenv("KAFKA_GROUP", "test_unit")
	os.Setenv("KAFKA_TOPICS", "test_topic")
	os.Setenv("KAFKA_ASSIGNOR", "range")
	os.Setenv("KAFKA_AUTHENTICATION_TYPE", "ssl")
	os.Setenv("KAFKA_AUTHENTICATION_CA", "test_authentication_ca")
	os.Setenv("KAFKA_AUTHENTICATION_CERTIFICATE", "test_authentication_certificate")
	os.Setenv("KAFKA_AUTHENTICATION_KEY", "test_authentication_key")

	kafkaConfig, err := GetKafka()

	assert.Empty(t, err)
	assert.Equal(t, "range", kafkaConfig.Assignor)
	assert.Equal(t, "test_topic", kafkaConfig.Topic)
	assert.Equal(t, "test_unit", kafkaConfig.ConsumerGroup)
	assert.Equal(t, "2.1.1", kafkaConfig.Version)
	assert.Equal(t, "test_broker", kafkaConfig.Brokers)
	assert.Equal(t, "ssl", kafkaConfig.AuthType)
	assert.Equal(t, "test_authentication_ca", kafkaConfig.AuthCa)
	assert.Equal(t, "test_authentication_certificate", kafkaConfig.AuthCertificate)
	assert.Equal(t, "test_authentication_key", kafkaConfig.AuthKey)
}

func TestGetKafkaConfigWithInvalidSSL(t *testing.T) {
	os.Setenv("KAFKA_BROKERS", "test_broker")
	os.Setenv("KAFKA_VERSION", "2.1.1")
	os.Setenv("KAFKA_GROUP", "test_unit")
	os.Setenv("KAFKA_TOPICS", "test_topic")
	os.Setenv("KAFKA_ASSIGNOR", "range")
	os.Setenv("KAFKA_AUTHENTICATION_TYPE", "ssl")
	os.Setenv("KAFKA_AUTHENTICATION_CA", "")
	os.Setenv("KAFKA_AUTHENTICATION_CERTIFICATE", "")
	os.Setenv("KAFKA_AUTHENTICATION_KEY", "")

	kafkaConfig, err := GetKafka()

	assert.Error(t, err)
	assert.Equal(t, "range", kafkaConfig.Assignor)
	assert.Equal(t, "test_topic", kafkaConfig.Topic)
	assert.Equal(t, "test_unit", kafkaConfig.ConsumerGroup)
	assert.Equal(t, "2.1.1", kafkaConfig.Version)
	assert.Equal(t, "test_broker", kafkaConfig.Brokers)
	assert.Equal(t, "ssl", kafkaConfig.AuthType)
}
