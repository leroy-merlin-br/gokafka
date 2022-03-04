package consumer

import (
	"github.com/Shopify/sarama"
	"github.com/golang/mock/gomock"
	"github.com/leroy-merlin-br/gokafka/consumer/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSaleOrderFromAvroRecordWithFewAttributes(t *testing.T) {
	// Set
	controller := gomock.NewController(t)
	claim := mocks.NewMockConsumerGroupClaim(controller)
	session := mocks.NewMockConsumerGroupSession(controller)

	action := func(message *sarama.ConsumerMessage) error {
		return nil
	}

	consumer := Consumer{
		Ready:  make(chan bool),
		Action: action,
	}

	messages := make(chan *sarama.ConsumerMessage)
	close(messages)

	// Expectations
	claim.EXPECT().Messages().Return(messages).Times(1)

	// Actions
	error := consumer.ConsumeClaim(session, claim)

	// Assertions
	assert.Equal(t, nil, error)
}
