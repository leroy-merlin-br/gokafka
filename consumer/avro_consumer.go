package consumer

import (
	"bytes"
	"github.com/Shopify/sarama"
	"github.com/linkedin/goavro"
	"github.com/pkg/errors"
	"reflect"
)

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *AvroConsumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.Ready)

	return nil
}

func (consumer *AvroConsumer) AvroDecode(message *sarama.ConsumerMessage) (*goavro.Record, error) {
	bb := bytes.NewBuffer(message.Value[5:])
	decoded, _ := consumer.Codec.Decode(bb)
	record, ok := decoded.(*goavro.Record)
	if !ok {
		return nil, errors.New("Type: " + reflect.TypeOf(decoded).String() + "is not a valid Record.")
	}

	return record, nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *AvroConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (consumer *AvroConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		record, err := consumer.AvroDecode(message)
		if err != nil {
			return err
		}

		err = consumer.Action(record)
		if err != nil {
			return err
		}

		session.MarkMessage(message, "")
	}

	return nil
}
