package main

import (
	"github.com/NygmaC/streaming-video/fileSentry/config"
	"github.com/NygmaC/streaming-video/fileSentry/internal/cache"
	"github.com/NygmaC/streaming-video/fileSentry/internal/consumer"
	"github.com/NygmaC/streaming-video/fileSentry/internal/dispatcher"
	"github.com/NygmaC/streaming-video/fileSentry/internal/handler"
	"github.com/NygmaC/streaming-video/fileSentry/internal/observer"
	"github.com/NygmaC/streaming-video/fileSentry/internal/producer"
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
