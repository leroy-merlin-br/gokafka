package consumer

import (
	"github.com/Shopify/sarama"
	"github.com/linkedin/goavro"
)

type AvroAction func(record *goavro.Record) error

type Action func(record *sarama.ConsumerMessage) error

type ConsumerInterface interface {
	// Setup is run at the beginning of a new session, before ConsumeClaim.
	Setup(sarama.ConsumerGroupSession) error

	// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
	// but before the offsets are committed for the very last time.
	Cleanup(sarama.ConsumerGroupSession) error

	// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
	// Once the Messages() channel is closed, the Handler must finish its processing
	// loop and exit.
	ConsumeClaim(sarama.ConsumerGroupSession, sarama.ConsumerGroupClaim) error

	IsReady() chan bool
	
	SetReady(ready chan bool)
}
