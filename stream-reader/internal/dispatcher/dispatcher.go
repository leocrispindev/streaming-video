package dispatcher

import (
	"encoding/json"

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

	for video.Read() {

		v := model_commons.MessageProccess{
			Value:  video.FrameBuffer(),
			Index:  index,
			Action: 0,
		}

		json, _ := json.Marshal(v)

		p.Produce(createMessage(proccess.TopicName, json), nil)

		index += 1
	}

	finish(p, proccess)
}

func finish(p *kafka.Producer, proccess model.Proccess) {

	v := model_commons.MessageProccess{
		Value:  []byte{},
		Index:  0,
		Action: 1,
	}

	json, _ := json.Marshal(v)

	p.Produce(createMessage(proccess.TopicName, json), nil)
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
