package producer

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
)

type ProducerInterface interface {
	SendMessage(topic string, value []byte) error
	SendMessageToPartition(topic string, value []byte, partition int32) error
}

type Producer struct {
	ProducerImpl sarama.SyncProducer
}

func (p *Producer) SendMessage(topic string, value []byte) error {

	_, _, err := p.ProducerImpl.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(value),
	})

	return err
}

func (p *Producer) SendMessageToPartition(topic string, value []byte, part int32) error {

	_, _, err := p.ProducerImpl.SendMessage(&sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.ByteEncoder(value),
		Partition: part,
	})

	return err
}

func getProducerConfig() *sarama.Config {
	producerConfig := sarama.NewConfig()

	producerConfig.Version = sarama.V2_2_0_0 // versão do Kafka
	producerConfig.Producer.Return.Successes = true
	producerConfig.Producer.MaxMessageBytes = 3000000

	return producerConfig
}

func CreateSyncProducer() Producer {

	producer, err := sarama.NewSyncProducer([]string{os.Getenv("KAFKA_HOST")}, getProducerConfig())
	if err != nil {
		log.Fatalln("Erro ao criar produtor:", err)
	}

	p := Producer{
		ProducerImpl: producer,
	}

	return p
}
