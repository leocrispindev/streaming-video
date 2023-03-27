package producer

import (
	"github.com/NygmaC/streamming-video/stream-go-commons/pkg/broker/producer"
)

var producerLocal producer.Producer

func Init() {

	producerLocal = producer.CreateSyncProducer()
}

func GetProducer() producer.Producer {
	if (producerLocal == producer.Producer{}) {
		Init()
	}

	return producerLocal
}
