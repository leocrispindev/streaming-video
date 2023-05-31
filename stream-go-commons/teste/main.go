package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/leocrispindev/streamming-video/stream-go-commons/pkg/broker/consumer"
	"github.com/leocrispindev/streamming-video/stream-go-commons/pkg/broker/producer"

	vidio "github.com/AlexEidt/Vidio"
)

func main() {

	os.Setenv("KAFKA_HOST", "localhost:9092")

	p := producer.CreateSyncProducer()

	video, err := vidio.NewVideo("/home/leonardo/Vídeos/projeto/video1.mp4")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", video.Frames())

	index := 0

	go func(index int) {
		for video.Read() {

			err = p.SendMessage("teste-leo", video.FrameBuffer())

			if err != nil {
				log.Fatalln(err)
			}

			index++

		}
	}(index)

	c := consumer.CreateConsumer("", "teste-leo")

	c.ReadMessage(HandleConsumer)

}

func HandleConsumer(msgs <-chan *sarama.ConsumerMessage) {

	for m := range msgs {
		fmt.Printf("Mensagem recebida %s", string(m.Key))

	}
}
