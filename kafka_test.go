package gokafka

import (
	"github.com/golang/mock/gomock"
	"github.com/leroy-merlin-br/gokafka/mocks"
	"testing"
)

func TestGetSaleOrderFromAvroRecordWithFewAttributes(t *testing.T) {
	// Set
	controller := gomock.NewController(t)
	consumer := mocks.NewMockConsumerInterface(controller)

	Handle(consumer)
}
