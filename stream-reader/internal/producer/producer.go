package producer

import (
	"os"

	"github.com/NygmaC/streamming-video/stream-go-commons/pkg/broker/producer"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var producerLocal *kafka.Producer

func Init() {

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	producerLocal = producer.CreateProducer(hostname)
}

func GetProducer() *kafka.Producer {
	if producerLocal == nil {
		Init()
	}

	return producerLocal
}
