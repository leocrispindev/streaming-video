package main

import (
	"github.com/NygmaC/streamming-video/stream-writter/config"
	"github.com/NygmaC/streamming-video/stream-writter/internal/broker/consumer"
)

func main() {
	config.Init()
	consumer.Init()
}
