package producer

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func CreateProducer(clientID string) *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"max.message.bytes": "2000000",
	})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	return p

}
