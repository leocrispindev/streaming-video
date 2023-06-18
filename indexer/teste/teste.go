package main

import (
	"github.com/leocrispindev/streaming-video/indexer/internal/indexer"
	"github.com/leocrispindev/streaming-video/indexer/internal/model"
)

func main() {
	indexer.Init()

	// Exemplo de criação de um objeto Video
	video := model.Video{
		ID:         4,
		Titulo:     "video do cachorro",
		Descricao:  "Descrição do Vídeo",
		Category:   2,
		Duracao:    120.5,
		Indexless:  true,
		Repository: "video",
	}

	indexer.Index(video)

	data := model.DeleteVideo{
		ID:         1,
		Repository: "video",
	}
	indexer.Delete(data)
}
