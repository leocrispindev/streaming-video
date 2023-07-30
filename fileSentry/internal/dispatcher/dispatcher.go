package dispatcher

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/leocrispindev/streaming-video/fileSentry/internal/cache"
	"github.com/leocrispindev/streaming-video/fileSentry/internal/producer"
)

var fileChannel chan string

func Init() chan string {
	fileChannel = make(chan string)

	go dispatch()

	fmt.Println("Dispatcher OK")

	return fileChannel
}

func dispatch() {

	re := regexp.MustCompile(`\b(\w+)\.\w+$`)

	for file := range fileChannel {

		match := re.FindStringSubmatch(file)
		if match == nil {
			//LOG
			return
		}

		fileName := match[0]

		videoID := strings.Split(fileName, "_")[0]
		//Busca no cache para qual topico enviar
		//Envia a mensagem informando qual o arrquuivo criado
		topic := cache.Get(videoID)

		log.Println("Send segment to topic: " + topic)

		msg, err := createNessage(videoID, fileName)
		if err != nil {
			println("Error on create message to topic: " + topic + ", Error: " + err.Error())
			continue
		}

		producer.ProducerLocal.SendMessage(topic, msg)

	}

}

func createNessage(videoID string, filePath string) ([]byte, error) {
	//message := model.NotifyMessage{VideoID: videoID, FileName: filePath}

	//return json.Marshal(message)
	return []byte(videoID + "/" + filePath), nil

}
