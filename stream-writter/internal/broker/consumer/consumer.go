package consumer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/NygmaC/streamming-video/stream-go-commons/pkg/broker/consumer"
	model_commons "github.com/NygmaC/streamming-video/stream-go-commons/pkg/model"
	"github.com/NygmaC/streamming-video/stream-writter/internal/writter"
	"github.com/Shopify/sarama"
)

func Init() {
	notifyTopic := os.Getenv("KAFKA_NOTIFY_WRITTER_TOPIC")

	notifyConsumer := consumer.CreateConsumer("", notifyTopic)

	// Consumidor responsavel por iniciar o processo para consumo dos frames
	notifyConsumer.ReadMessage(handleMessage)

}

func handleMessage(channelMsg <-chan *sarama.ConsumerMessage) {

	for msg := range channelMsg {

		notify := model_commons.NotifyProccess{}

		if err := parse(msg.Value, &notify); err != nil {
			fmt.Println("Error parse notify: " + string(msg.Key))
			continue
		}

		writter.StartProccess(notify)

		// TODO Chamar o responsavel por lidar com os processos criados

	}
}

func parse(value []byte, n *model_commons.NotifyProccess) error {

	return json.Unmarshal(value, n)

}
