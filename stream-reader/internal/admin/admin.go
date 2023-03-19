package admin

import (
	"context"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var kAdmin *kafka.AdminClient

func Init() {
	createAdminFromConsumer()
}

func createAdminFromConsumer() {

	kAdminResult, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_HOST"),
	})

	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
		os.Exit(1)
	}

	kAdmin = kAdminResult

	println("Admin OK")
}

func CreateTopic(ctx context.Context, numberParititon int, topicName string) (string, error) {

	topic := kafka.TopicSpecification{
		Topic:         topicName,
		NumPartitions: numberParititon,
	}

	result, err := kAdmin.CreateTopics(ctx, []kafka.TopicSpecification{topic})

	if err != nil {
		return "", err
	}

	return result[0].Topic, nil
}
