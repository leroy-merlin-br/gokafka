# GOKafka

> Easy and flexible Kafka Library for GO.

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)

- [Introduction](#introduction)
- [Requirements](#requirements)
- [Installation](#installation)
- [Avro Schema Quick Usage Guide](#avro-schema-quick-usage-guide)
- [License](#license)


<a name="introduction"></a>
## Introduction

**GOKafka** provides a simple, straight-forward implementation for working with Kafka inside GO applications.

<a name="requirements"></a>
## Requirements

- GO >= 1.17

<a name="installation"></a>
## Installation

You can install the library via GO get:

```
$ go get github.com/leroy-merlin-br/gokafka
```

And add the env`s for with:

```
KAFKA_BROKERS="example.southamerica-east1.gcp.confluent.cloud:9092"
KAFKA_VERSION=2.1.1
KAFKA_GROUP="kafka-group-example"
KAFKA_TOPICS="EXAMPLE-TOPIC-V1"
KAFKA_ASSIGNOR="range"
KAFKA_AUTHENTICATION_TYPE="sasl_ssl"

# to use sasl_ssl authentication
KAFKA_USERNAME=
KAFKA_PASSWORD=

# to use ssl authentication
KAFKA_AUTHENTICATION_TYPE=
KAFKA_AUTHENTICATION_CA=
KAFKA_AUTHENTICATION_KEY=
KAFKA_AUTHENTICATION_CERTIFICATE=
```

If you will use Avro Schema, you must add this env`s too:

```
AVRO_SCHEMA_URL="https://sr-southamerica-east1.streaming.data.cloud"
AVRO_SCHEMA_USERNAME=
AVRO_SCHEMA_PASSWORD=
```

## Avro Schema Quick Usage Guide

Create kafka_adapter.go
```
import (
	"github.com/leroy-merlin-br/gokafka"
	"github.com/linkedin/goavro"
	"github.com/rs/zerolog/log"
	"lmbr/saleorder/domain"
)

type KafkaAdapter struct {
}

type User struct {
	Id   string
	Name string
}

// Adapts a message action to Kafka Action
func (adapter *KafkaAdapter) Receive(saleOrderAction domain.Action) error {
	action := func(record *goavro.Record) error {
		decodedRecord := adapter.Transform(record)

		log.Print(decodedRecord)

		return nil
	}

	consumer := gokafka.Consumer{
		Ready:  make(chan bool),
		Action: action,
	}

	return gokafka.Handle(consumer)
}

// Adapts a message (kafka Record in this case) to a User
func (adapter *KafkaAdapter) Transform(record *goavro.Record) User {
	id, _ := record.Get("id")
	name, _ := record.Get("name")

	user := User{
		Id:   id.(string),
		Name: name.(string),
	}

	return user
}
```

Create worker.go
```
func main() {
    messageService := KafkaAdapter{}
    err := messageService.Receive(saveAction)
    if err != nil {
        log.Error().Err(err).Msg("Receive handler error occurs.")
        return err
    }
    
    return nil
}
```

<a name="license"></a>
## License

**GOKafka** is free software distributed under the terms of the [MIT license](http://opensource.org/licenses/MIT)

<a name="additional_information"></a>
## Additional information

**GOKafka** was proudly built by the [Leroy Merlin Brazil](https://github.com/leroy-merlin-br) team. [See all the contributors](https://github.com/leroy-merlin-br/gokafka/graphs/contributors).
