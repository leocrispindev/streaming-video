package producer

import (
	"github.com/NygmaC/streamming-video/stream-go-commons/pkg/broker/producer"
)

var ProducerLocal producer.Producer

func Init() {

	ProducerLocal = producer.CreateSyncProducer()
}
