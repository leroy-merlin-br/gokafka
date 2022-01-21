package config

import (
	"errors"
	"os"
)

type AvroConfig struct {
	Url      string
	Username string
	Password string
}

// Creating error vars like this
// We'll be able to use it to check what kind of error without
// reading the content of error.
// e.g: if (err == config.AVRO_SCHEMA_URL)
var (
	invalidUrl = errors.New("no Avro url AvroConfig.Url defined, please set the AVRO_SCHEMA_URL env")
)

func GetAvro() (config AvroConfig, err error) {
	avroConfig := AvroConfig{
		Url:      os.Getenv("AVRO_SCHEMA_URL"),
		Username: os.Getenv("AVRO_SCHEMA_USERNAME"),
		Password: os.Getenv("AVRO_SCHEMA_PASSWORD"),
	}

	if len(avroConfig.Url) == 0 {
		return avroConfig, invalidUrl
	}

	return avroConfig, nil
}
