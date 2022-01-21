package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAvroConfig(t *testing.T) {
	os.Setenv("AVRO_SCHEMA_URL", "http://schema-registry")
	os.Setenv("AVRO_SCHEMA_USERNAME", "test_user")
	os.Setenv("AVRO_SCHEMA_PASSWORD", "test_password")

	avroConfig, err := GetAvro()

	assert.Empty(t, err)
	assert.Equal(t, "http://schema-registry", avroConfig.Url)
	assert.Equal(t, "test_user", avroConfig.Username)
	assert.Equal(t, "test_password", avroConfig.Password)
}

func TestGetAvroErrorOnInvalidConfig(t *testing.T) {
	os.Setenv("AVRO_SCHEMA_URL", "")
	os.Setenv("AVRO_SCHEMA_USERNAME", "test_user")
	os.Setenv("AVRO_SCHEMA_PASSWORD", "test_password")

	_, err := GetAvro()
	assert.Error(t, err)
}
