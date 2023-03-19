package dispatcher

import (
	"encoding/json"
	"fmt"

	vidio "github.com/AlexEidt/Vidio"
	model_commons "github.com/NygmaC/streamming-video/stream-go-commons/pkg/model"
	"github.com/NygmaC/streamming-video/stream-reader/internal/model"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var packageSize int

func Init() {
}

// Action
// 0 - Add messages in buffer on writter
// 1 - The dispatcher is over

func Start(p *kafka.Producer, video *vidio.Video, proccess model.Proccess) {

	index := 0

	fmt.Println("Proccess started for topic: " + proccess.TopicName)

	for video.Read() {

		v := model_commons.MessageProccess{
			Value:  video.FrameBuffer(),
			Index:  index,
			Action: 0,
		}

		json, _ := json.Marshal(v)

		fmt.Println(len(v.Value))

		err := p.Produce(createMessage(proccess.TopicName, json), nil)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Frame send to topic: " + proccess.TopicName)

		index += 1
	}

	finish(p, proccess, index)
}

func finish(p *kafka.Producer, proccess model.Proccess, index int) {

	v := model_commons.MessageProccess{
		Value:  []byte{},
		Index:  index,
		Action: 1,
	}

	json, _ := json.Marshal(v)

	p.Produce(createMessage(proccess.TopicName, json), nil)

	fmt.Println("Proccess is finished: " + proccess.TopicName)
}

func createMessage(topic string, value []byte) *kafka.Message {

	kM := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic, Partition: kafka.PartitionAny,
		},
		Value: value,
	}

	return &kM
}
