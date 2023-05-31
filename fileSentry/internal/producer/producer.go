package producer

import (
	"fmt"

	"github.com/leocrispindev/streamming-video/stream-go-commons/pkg/broker/producer"
)

var ProducerLocal producer.Producer

func Init() {

	ProducerLocal = producer.CreateSyncProducer()
	fmt.Println("Producer OK")
}
