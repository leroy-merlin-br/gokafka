package consumer

import (
	"github.com/Shopify/sarama"
)

type Consumer struct {
	Ready  chan bool
	Action Action
}

func (consumer *Consumer) IsReady() chan bool {
	return consumer.Ready
}

func (consumer *Consumer) SetReady(ready chan bool) {
	consumer.Ready = ready
}

func (consumer *Consumer) getAction() Action {
	return consumer.Action
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
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		err := consumer.Action(message)
		if err != nil {
			return err
		}

		session.MarkMessage(message, "")
	}

	return nil
}
