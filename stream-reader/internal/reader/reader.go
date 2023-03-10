package reader

import (
	"log"
	"os"

	vidio "github.com/AlexEidt/Vidio"
	"github.com/NygmaC/streamming-video/stream-reader/internal/model"
)

// Struct for local storage configuration
type storageLocal struct {
	basePath string
}

// Define reader interface
type Reader interface {
	read(filename string) (*vidio.Video, error)
}

func (s storageLocal) read(filename string) (*vidio.Video, error) {
	return vidio.NewVideo(s.basePath + filename)
}

var reader Reader

func init() {

	reader = storageLocal{os.Getenv("STORAGE_LOCAL_BASE_PATH")}
	//Define o reader

}

// Monta o pacode de frames do video, enviar para o producer
func Proccess(p model.Proccess) {

	video, err := reader.read(p.VideoName)

	if err != nil {
		log.Fatalf("Error", err)
	}

	for video.Read() {
		// Todo 
	}

}
