package reader

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	vidio "github.com/AlexEidt/Vidio"
	commons "github.com/NygmaC/streamming-video/stream-go-commons/pkg/model"
	adminInternal "github.com/NygmaC/streamming-video/stream-reader/internal/admin"
	"github.com/NygmaC/streamming-video/stream-reader/internal/dispatcher"
	"github.com/NygmaC/streamming-video/stream-reader/internal/model"
	prodInternal "github.com/NygmaC/streamming-video/stream-reader/internal/producer"
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
	return vidio.NewVideo(s.BasePath + "/" + filename)
}

var reader Reader

var notifyWritterTopic string

func Init() {

	// TODO definir o storage via property
	reader = storageLocal{
		BasePath: os.Getenv("STORAGE_LOCAL_BASE_PATH"),
	}

	fmt.Println("Reader OK")

	notifyWritterTopic = os.Getenv("KAFKA_NOTIFY_WRITTER_TOPIC")

}

// Monta o paco de de frames do video, enviar para o producer
func Proccess(p model.Proccess) {

	// TODO verificar existencia do topico para o processo
	// se não existir, criar
	// Após criar notificar o writter

	topicName := "proccess-stream-" + p.VideoName

	err := adminInternal.GetAdminInstance().CreateTopic(topicName, 1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created topic: " + topicName)

	p.TopicName = topicName

	go notifyWritter(p)
	video, err := readFile(p)

	if err != nil {
		println("Error on read file")
		return
	}

	go dispatcher.Start(prodInternal.ProducerLocal, video, p)
}

// Envia uma notificação para o Writter de que um processo vai iniciar
// Writter deverá criar um consumer para o processo
func notifyWritter(p model.Proccess) {

	// TODO passar o nome do topico do processo que deverá ser iniciado
	msg := commons.NotifyProccess{
		ProccessID: fmt.Sprint(p.Id),
		VideoName:  p.VideoName,
		Action:     0,
		TopicName:  p.TopicName,
	}

	msgJson, _ := json.Marshal(msg)

	prodInternal.ProducerLocal.SendMessage(notifyWritterTopic, msgJson)
}

// Faz a leitura do arquivo
func readFile(p model.Proccess) (*vidio.Video, error) {
	video, err := reader.read(p.VideoName)

	if err != nil {
		log.Println("Error", err)
		return nil, err
	}

	fmt.Printf("Video read: %s", video.FileName())

	return video, nil
}
