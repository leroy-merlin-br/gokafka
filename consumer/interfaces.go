package consumer

import (
	"github.com/Shopify/sarama"
	"github.com/linkedin/goavro"
)

type AvroAction func(record *goavro.Record) error

type Action func(record *sarama.ConsumerMessage) error

type Consumer struct {
	Ready  chan bool
	Action Action
}

type AvroConsumer struct {
	Ready  chan bool
	Action AvroAction
	Codec  goavro.Codec
}
