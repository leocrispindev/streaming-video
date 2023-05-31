package main

import (
	"github.com/leocrispindev/streaming-video/fileSentry/config"
	"github.com/leocrispindev/streaming-video/fileSentry/internal/consumer"
)

func main() {
	config.Init()
	consumer.Init()
}
