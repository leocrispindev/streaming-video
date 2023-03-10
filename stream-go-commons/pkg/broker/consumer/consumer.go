package consumer

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func CreateConsumer(groupID string, topics []string) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_HOST"),
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = c.SubscribeTopics(topics, nil)

	if err != nil {
		panic(err)
	}
}
