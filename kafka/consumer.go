package kafka

import (
	"context"
	"encoding/json"

	"github.com/Shopify/sarama"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog/log"

	"lmbr/merchant/merchant"
)

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	Ready chan bool
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
	ctx := context.Background()
	option := merchant.MerchantConfig{
		Path: "/app/merchant/service-account.json",
	}
	service, err := merchant.NewMerchant(ctx, option)
	if err != nil {
		panic(err)
	}

	log.Debug().Msg("Start consuming messages to send to merchant center.")

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	for message := range claim.Messages() {
		entries, err := getEntries(message)
		if err != nil {
			log.Error().Err(err).Msg("Error when decoding message to json:")
		}

		// Currently we are receiving only one message per kafka message
		// But the schema was designed thinking that we could add more messages
		// if neccessary, wo we wil loop on this entries, just in case.
		for _, entry := range entries {
			if "delete" == entry.Method {
				service.DeleteRegionalInventory(ctx, entry.ProductId)

				continue
			}

			service.SendRegionalInventory(ctx, &entry)
		}

		// Commit message and mark it as consumed.
		session.MarkMessage(message, "")
	}

	return nil
}

// Use json.Unmarshal to decode the message into a json object
// If fails, return error.
func getEntries(message *sarama.ConsumerMessage) ([]merchant.RegionalEntry, error) {
	var entries merchant.RegionalInventoryRecord
	if err := json.Unmarshal(message.Value, &entries); err != nil {
		return nil, err
	}

	return entries.Entries, nil
}
