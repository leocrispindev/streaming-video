package producer

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func CreateProducer(clientID string) *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_HOST"),
		"client.id":         clientID,
		"acks":              "all",
		"max.request.size":  3000000})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		panic(err)
	}

	return p

}
