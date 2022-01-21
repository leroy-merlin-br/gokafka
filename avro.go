package gokafka

import (
	"github.com/linkedin/goavro"
	"github.com/riferrei/srclient"
)

func Codec() (goavro.Codec, error) {
	kafkaConfig, err := GetKafka()
	if err != nil {
		return nil, err
	}

	avroConfig, err := GetAvro()
	if err != nil {
		return nil, err
	}

	schemaRegistryClient := srclient.CreateSchemaRegistryClient(avroConfig.Url)
	schemaRegistryClient.SetCredentials(avroConfig.Username, avroConfig.Password)
	avroSchema, err := schemaRegistryClient.GetLatestSchema(kafkaConfig.Topic + "-value")
	if err != nil {
		return nil, err
	}

	codec, err := goavro.NewCodec(avroSchema.Schema())
	if err != nil {
		return nil, err
	}

	return codec, err
}
