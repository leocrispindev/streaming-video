package dispatcher

import (
	"encoding/json"
	"fmt"

	vidio "github.com/AlexEidt/Vidio"
	prodCommons "github.com/NygmaC/streamming-video/stream-go-commons/pkg/broker/producer"
	model_commons "github.com/NygmaC/streamming-video/stream-go-commons/pkg/model"
	"github.com/NygmaC/streamming-video/stream-reader/internal/model"
)

func Init() {
}

// Action
// 0 - Add messages in buffer on writter
// 1 - The dispatcher is over

func Start(p prodCommons.Producer, video *vidio.Video, proccess model.Proccess) {

	index := 0

	fmt.Println("Proccess started for topic: " + proccess.TopicName)

	framesBySegment := video.Frames() / proccess.Segments

	for partitionNumber := 1; partitionNumber < proccess.Segments; partitionNumber++ {

		index = 0

		for index < framesBySegment {

			if !video.Read() {
				break
			}

			v := model_commons.MessageProccess{
				Value:  video.FrameBuffer(),
				Index:  index,
				Action: 0,
			}

			json, _ := json.Marshal(v)

			//Envia para a partição definida per partitionNumber
			err := p.SendMessageToPartition(proccess.TopicName, json, int32(partitionNumber))

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(fmt.Printf("partition proccess: %d, to Topic: %s", partitionNumber, proccess.TopicName))

			index += 1
		}

		finishPartition(p, proccess, partitionNumber)
	}

	finish(p, proccess, index)
}

func finishPartition(p prodCommons.Producer, proccess model.Proccess, partition int) {
	v := model_commons.MessageProccess{
		Value:  []byte{},
		Index:  0,
		Action: 1,
	}

	value, _ := json.Marshal(v)

	p.SendMessageToPartition(proccess.TopicName, value, int32(partition))

	fmt.Println(fmt.Printf("Finish partition proccess: %d, to Topic: %s", partition, proccess.TopicName))
}

func finish(p prodCommons.Producer, proccess model.Proccess, index int) {

	v := model_commons.MessageProccess{
		Value:  []byte{},
		Index:  index,
		Action: 1,
	}

	json, _ := json.Marshal(v)

	p.SendMessage(proccess.TopicName, json)

	fmt.Println("Proccess is finished: " + proccess.TopicName)
}
