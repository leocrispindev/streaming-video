package reader

import (
	"fmt"
	"log"
	"os"

	vidio "github.com/AlexEidt/Vidio"
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

func Init() {

	// TODO definir o storage via property
	reader = storageLocal{
		BasePath: os.Getenv("STORAGE_LOCAL_BASE_PATH"),
	}

	fmt.Println("Reader OK")
	//Define o reader
	prod = producer.GetProducer()

}

// Monta o paco de de frames do video, enviar para o producer
func Proccess(p model.Proccess) {

	videoname := p.VideoName

	fmt.Println(videoname)

	notifyWritter(p)
	read(videoname)
}

// Envia uma notificação para o Writter de que um processo vai iniciar
//Writter deverá criar um consumer para o processo
func notifyWritter(p model.Proccess) {

	topic := "myTopic"

	prod.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic, Partition: kafka.PartitionAny,
		},
	}, nil)

}

// Faz a leitura do arquivo, e envia os frames para o writter
func read(videoname string) {
	video, err := reader.read(videoname)

	if err != nil {
		log.Println("Error", err)
		return
	}

	for video.Read() {

		// O pacote é definido atraves da quantidade de FPS, 1s tem X fps
		// um pacote deve ter no minimo 2 segundos
		//TODO envia os buffers para o writter do processo
	}
}
