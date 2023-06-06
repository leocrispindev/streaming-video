package main

import (
	"github.com/leocrispindev/streaming-video/indexer/config"
	"github.com/leocrispindev/streaming-video/indexer/internal/consumer"
	"github.com/leocrispindev/streaming-video/indexer/internal/indexer"
)

func main() {
	config.Init()
	indexer.Init()
	consumer.Init()
}
