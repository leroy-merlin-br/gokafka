package consumer

import (
	"errors"
	"github.com/Shopify/sarama"
	"github.com/golang/mock/gomock"
	"github.com/leroy-merlin-br/gokafka/consumer/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func addMessage(message *sarama.ConsumerMessage, c chan *sarama.ConsumerMessage) {
	c <- message
}

func TestGetSaleOrderFromAvroRecordWithFewAttributes(t *testing.T) {
	// Set
	controller := gomock.NewController(t)
	claim := mocks.NewMockConsumerGroupClaim(controller)
	session := mocks.NewMockConsumerGroupSession(controller)

	action := func(message *sarama.ConsumerMessage) error {
		if message.Partition != 1 {
			return errors.New("Invalid Partition")
		}

		return nil
	}

	consumer := Consumer{
		Ready:  make(chan bool),
		Action: action,
	}

	messages := make(chan *sarama.ConsumerMessage)
	message := &sarama.ConsumerMessage{
		Partition: 1,
	}

	messageB := &sarama.ConsumerMessage{
		Partition: 2,
	}

	go addMessage(message, messages)
	go addMessage(messageB, messages)

	// Expectations
	claim.EXPECT().Messages().Return(messages)

	session.EXPECT().MarkMessage(message, "")

	// Actions
	error := consumer.ConsumeClaim(session, claim)

	// Assertions
	assert.Equal(t, errors.New("Invalid Partition"), error)
}
