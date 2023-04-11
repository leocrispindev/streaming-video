package writter

import (
	"encoding/json"
	"fmt"

	vidio "github.com/AlexEidt/Vidio"
	"github.com/NygmaC/streamming-video/stream-go-commons/pkg/broker/consumer"
	model_commons "github.com/NygmaC/streamming-video/stream-go-commons/pkg/model"
	"github.com/Shopify/sarama"
)

func StartProccess(notify model_commons.NotifyProccess) {

	seg := 1

	framesConsumer := consumer.CreateConsumer("", notify.TopicName)

	videoName := notify.VideoName

	fmt.Println("Proccess de leitura recebido para o topico: " + notify.TopicName)

	for seg <= notify.VideoOptions.Segments {

		//framesConsumer.ReadMessage(handleMessage)

		// Recebe a quantiade de segmentos e dispara um consumidor para cada partição
		// Cada consumidor vai receber um numero X de frames, é o valor já do segmento, só agrupar
		// e escrever
		go initPartitionConsumer(int32(seg), framesConsumer, notify.VideoOptions, videoName)
		seg++

	}

}

func initPartitionConsumer(partiton int32, consumer consumer.Consumer, videoOptions model_commons.VideoOptions, videoName string) {
	partition, err := consumer.ConsumerImpl.ConsumePartition(consumer.Topic, partiton, sarama.OffsetOldest)

	if err != nil {
		fmt.Sprintln("Erro on Read Message: ", err.Error())
		return

	}

	println(fmt.Sprintf("Iniciando consumo de partição %d, do topico %s", partition, consumer.Topic))

	options := vidio.Options{
		FPS:     videoOptions.Fps,
		Bitrate: videoOptions.Bitrate,
	}

	segmentName := fmt.Sprintf("%s-%s-%d.mp4", videoName, "segment", partiton)

	wSegment, _ := vidio.NewVideoWriter(segmentName, videoOptions.Width, videoOptions.Height, &options)

	for fr := range partition.Messages() {

		fp := parse(fr.Value)

		if fp.Action == 1 {
			break
		}

		wSegment.Write(fp.Value)

	}

	wSegment.Close()
	println(fmt.Sprintf("Encerrando escrita do arquivo:  %s", segmentName))

	partition.Close()
	println(fmt.Sprintf("Encerrando consumo de partição %d, do topico %s", partition, consumer.Topic))

}

func parse(value []byte) model_commons.MessageProccess {
	frameParse := model_commons.MessageProccess{}

	json.Unmarshal(value, &frameParse)

	return frameParse
}
