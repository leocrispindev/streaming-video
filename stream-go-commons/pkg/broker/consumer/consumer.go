package consumer

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
)

type ConsumerInterface interface {
	ReadMessage(handle func([]byte)) error
}

type Consumer struct {
	ConsumerImpl sarama.Consumer
	Topic        string
	Partition    int
}

func (c *Consumer) ReadMessage(f func(<-chan *sarama.ConsumerMessage)) {

	partition, err := c.ConsumerImpl.ConsumePartition(c.Topic, 0, sarama.OffsetOldest)

	if err != nil {
		log.Fatalf("Erro on Read Message %s", err.Error())

	}

	println(len(partition.Messages()))

	f(partition.Messages())

	//time.Sleep(1000)

}

func getConsumerConfig() *sarama.Config {
	consumerConfig := sarama.NewConfig()

	consumerConfig.Consumer.Return.Errors = true
	consumerConfig.Version = sarama.V2_2_0_0 // versão do Kafka

	return consumerConfig
}

func CreateConsumer(groupID string, topics string) Consumer {

	consumer, err := sarama.NewConsumer([]string{os.Getenv("KAFKA_HOST")}, getConsumerConfig())
	if err != nil {
		log.Fatalln("Erro ao criar consumidor:", err)
	}

	c := Consumer{
		ConsumerImpl: consumer,
		Topic:        topics,
	}

	return c

}
