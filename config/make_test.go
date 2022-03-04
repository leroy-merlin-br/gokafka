package config

import (
	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldMakeKafkaConfig(t *testing.T) {
	kafkaConfig := EnvKafkaConfig{
		Assignor: "range",
		Username: "KafkaUsername",
		Password: "KafkaPassword",
		AuthType: "sasl_ssl",
	}
	saramaConfig := sarama.NewConfig()

	_ = ConfigureSarama(kafkaConfig, saramaConfig)

	assert.Equal(t, sarama.BalanceStrategyRange, saramaConfig.Consumer.Group.Rebalance.Strategy)
	assert.Equal(t, "KafkaUsername", saramaConfig.Net.SASL.User)
	assert.Equal(t, "KafkaPassword", saramaConfig.Net.SASL.Password)
}
