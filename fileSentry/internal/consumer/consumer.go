package consumer

import (
	"encoding/json"
	"fmt"

	"github.com/NygmaC/streaming-video/fileSentry/internal/handler"
	"github.com/NygmaC/streaming-video/fileSentry/internal/model"
	"github.com/NygmaC/streamming-video/stream-go-commons/pkg/broker/consumer"
	"github.com/Shopify/sarama"
)

var proccessConsumer consumer.Consumer

func Init() {
	// {"videoName":"video2.mp4", "session":"aaaaaa", "connection": {}}

	proccessConsumer = consumer.CreateConsumer("", "stream-proccess")
	proccessConsumer.ReadMessage(handleMessage)

}

func handleMessage(msgs <-chan *sarama.ConsumerMessage) {
	fmt.Println("Consumer OK")

	for msg := range msgs {
		var streamProccess = model.Proccess{}

		err := parse(msg.Value, &streamProccess)

		if err != nil {
			fmt.Println(err)
			continue

		}

		go handler.Exec(streamProccess)
	}
}

func parse(value []byte, p *model.Proccess) error {
	fmt.Println(string(value))
	return json.Unmarshal(value, p)
}
