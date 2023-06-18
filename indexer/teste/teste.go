package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/leocrispindev/streaming-video/indexer/internal/indexer"
	"github.com/leocrispindev/streaming-video/indexer/internal/model"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	indexer.Init()

	video := &model.VideoInfo{
		ID:        4,
		Title:     "video do cachorro",
		Sinopse:   "Descrição do Vídeo",
		Category:  2,
		Duration:  1205,
		Indexless: true,
	}

	document := &model.Document{
		Action:     1,
		ID:         10,
		Key:        "video-index-id-8",
		Repository: "video",
		VideoInfo:  video,
	}

	body, err := json.Marshal(document.VideoInfo)

	if err != nil {
		fmt.Printf("Parse index error for KEY: %s\nMessage: %s\n", document.Key, err.Error())
		return
	}

	println(string(body))

	indexer.Index(*document)

	indexer.Delete(*document)

}
