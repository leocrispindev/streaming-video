package main

import (
	"net/http"
	"os"

	"github.com/leocrispindev/streaming-video/searcher/config"
	"github.com/leocrispindev/streaming-video/searcher/internal/controller"
)

func main() {
	config.Init()

	http.HandleFunc("/search", controller.Search)

	err := http.ListenAndServe(os.Getenv("HTTP_SERVER_PORT"), nil)

	if err != nil {
		panic(err)
	}
}
