package consumer

import (
	"encoding/json"
	"fmt"

	"github.com/NygmaC/streamming-video/stream-go-commons/pkg/broker/consumer"
	"github.com/NygmaC/streamming-video/stream-reader/internal/model"
	"github.com/NygmaC/streamming-video/stream-reader/internal/reader"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var proccessConsumer *kafka.Consumer

func Init() {
	// {"videoName":"teste1", "session":"aaaaaa", "connection": {}}
	proccessConsumer = consumer.CreateConsumer("process-group", []string{"video-stream-proccess"})

	start()
}

func GetConsumerInstance() *kafka.Consumer {
	return proccessConsumer
}

func start() {
	run()
}

func run() {
	fmt.Println("Consumer OK")

	for {
		ev := proccessConsumer.Poll(3000)

		switch e := ev.(type) {
		case *kafka.Message:

			fmt.Println("Get message")
			var streamProccess = model.Proccess{}

			err := parse(e.Value, &streamProccess)

			if err != nil {
				fmt.Println(err)
				return

			}

			go reader.Proccess(streamProccess)

		case kafka.Error:
			fmt.Println(e)
			//proccessConsumer.Close()
		}
	}
}

func parse(value []byte, p *model.Proccess) error {

	return json.Unmarshal(value, p)
}
