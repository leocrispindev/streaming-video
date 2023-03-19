package reader

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	vidio "github.com/AlexEidt/Vidio"
	commons "github.com/NygmaC/streamming-video/stream-go-commons/pkg/model"
	"github.com/NygmaC/streamming-video/stream-reader/internal/admin"
	"github.com/NygmaC/streamming-video/stream-reader/internal/dispatcher"
	"github.com/NygmaC/streamming-video/stream-reader/internal/model"
	"github.com/NygmaC/streamming-video/stream-reader/internal/producer"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// Struct for local storage configuration
type storageLocal struct {
	BasePath string
}

// Define reader interface
type Reader interface {
	read(filename string) (*vidio.Video, error)
}

func (s storageLocal) read(filename string) (*vidio.Video, error) {
	return vidio.NewVideo(s.BasePath + filename)
}

var reader Reader

var prod *kafka.Producer

var notifyWritterTopic string

func Init() {

	// TODO definir o storage via property
	reader = storageLocal{
		BasePath: os.Getenv("STORAGE_LOCAL_BASE_PATH"),
	}

	fmt.Println("Reader OK")

	notifyWritterTopic = os.Getenv("KAFKA_NOTIFY_WRITTER_TOPIC")

	//Define o reader
	prod = producer.GetProducer()

}

// Monta o paco de de frames do video, enviar para o producer
func Proccess(p model.Proccess) {

	// TODO verificar existencia do topico para o processo
	// se não existir, criar
	// Após criar notificar o writter
	topicName, err := admin.CreateTopic(context.Background(), 1, "proccess-stream-"+p.VideoName)

	if err != nil {
		log.Fatal(err)
	}

	p.TopicName = topicName

	go notifyWritter(p)
	video, err := readFile(p)

	go dispatcher.Start(prod, video, p)
}

// Envia uma notificação para o Writter de que um processo vai iniciar
//Writter deverá criar um consumer para o processo
func notifyWritter(p model.Proccess) {

	// TODO passar o nome do topico do processo que deverá ser iniciado
	msg := commons.NotifyProccess{
		ProccessID: fmt.Sprint(p.Id),
		VideoName:  p.VideoName,
		Action:     0,
	}

	msgJson, _ := json.Marshal(msg)

	prod.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &notifyWritterTopic, Partition: kafka.PartitionAny,
		},
		Value: []byte(msgJson),
	}, nil)

}

// Faz a leitura do arquivo
func readFile(p model.Proccess) (*vidio.Video, error) {
	video, err := reader.read(p.VideoName)

	if err != nil {
		log.Println("Error", err)
		return nil, err
	}

	return video, nil
}
