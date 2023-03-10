package main

import (
	"github.com/NygmaC/streamming-video/stream-reader/config"
	"github.com/NygmaC/streamming-video/stream-reader/pkg/server"
)

func main() {
	config.Init()
	server.Init()
}
