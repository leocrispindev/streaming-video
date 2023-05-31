package main

import (
	"github.com/leocrispindev/streaming-video/fileSentry/config"
	"github.com/leocrispindev/streaming-video/fileSentry/internal/cache"
	"github.com/leocrispindev/streaming-video/fileSentry/internal/consumer"
	"github.com/leocrispindev/streaming-video/fileSentry/internal/dispatcher"
	"github.com/leocrispindev/streaming-video/fileSentry/internal/handler"
	"github.com/leocrispindev/streaming-video/fileSentry/internal/observer"
	"github.com/leocrispindev/streaming-video/fileSentry/internal/producer"
)

func main() {
	config.Init()
	cache.Init()
	producer.Init()
	channel := dispatcher.Init()
	observer.Init(channel)
	handler.Init()
	consumer.Init()
}
