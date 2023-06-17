package consumer

import (
	"log"
	"os"
	"strconv"

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

	partition, err := c.ConsumerImpl.ConsumePartition(c.Topic, int32(c.Partition), sarama.OffsetOldest)

	if err != nil {
		log.Fatalf("Error on Read Message %s", err.Error())

	}

	println(len(partition.Messages()))

	f(partition.Messages())

	//time.Sleep(1000)

}

func getConsumerConfig() *sarama.Config {
	consumerConfig := sarama.NewConfig()

	consumerConfig.Consumer.Return.Errors = true
	consumerConfig.Version = sarama.V2_2_0_0 // versão do Kafka
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	consumerConfig.Consumer.Group.InstanceId = os.Getenv("KAFKA_GROUP_ID")
	return consumerConfig
}

func CreateConsumer(topics string) Consumer {

	consumer, err := sarama.NewConsumer([]string{os.Getenv("KAFKA_HOST")}, getConsumerConfig())
	if err != nil {
		log.Fatalln("Erro ao criar consumidor:", err)
	}

	partition, _ := strconv.Atoi(os.Getenv("KAFKA_CONSUMER_PARTITION"))

	c := Consumer{
		ConsumerImpl: consumer,
		Topic:        topics,
		Partition:    partition,
	}

	return c

}
