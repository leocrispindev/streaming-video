package handler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/leocrispindev/streaming-video/fileSentry/internal/cache"
	"github.com/leocrispindev/streaming-video/fileSentry/internal/generator"
	"github.com/leocrispindev/streaming-video/fileSentry/internal/model"
	"github.com/leocrispindev/streaming-video/fileSentry/internal/observer"
	utils "github.com/leocrispindev/streaming-video/fileSentry/internal/util"
)

var fileGenerator *generator.Generator

var pathFileSource string

func Init() {
	pathFileSource = os.Getenv("FILE_SOURCE_PATH")

	fileGenerator = generator.New()
}

func Exec(streamProccess model.Proccess) {
	videoName := streamProccess.VideoName

	//Path para o video original
	filePath := fmt.Sprintf("%s/%s.%s", pathFileSource, videoName, streamProccess.Extension)

	exportPath := filepath.Join(os.Getenv("FILE_STORAGE"), videoName)

	err := utils.CreatePath(exportPath)
	if err != nil {
		log.Println("Error on create path: " + err.Error())
	}

	//Add o path para o observer olhar para os paths de export
	observer.Add(exportPath)

	//Add no cache
	cache.Set(videoName, streamProccess.TopicName)

	cmdString, err := fileGenerator.Exec(filePath, exportPath, videoName)

	if err != nil {
		log.Println("Error on execute command: " + cmdString)
		observer.Remove(exportPath)

	}

}
