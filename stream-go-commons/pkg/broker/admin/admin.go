package admin

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
)

type AdminInterface interface {
	CreateTopic(name string, numPartitions int, numReplication int) error
	DeleteTopic(name string) error
}

type Admin struct {
	AdminImpl sarama.ClusterAdmin
}

var a Admin

func (a *Admin) CreateTopic(name string, numPartitions int, numReplication int) error {
	err := a.AdminImpl.CreateTopic(name, &sarama.TopicDetail{
		NumPartitions:     int32(numPartitions),
		ReplicationFactor: int16(numReplication),
	}, false)

	return err
}

func (a *Admin) DeleteTopic(name string) error {
	err := a.AdminImpl.DeleteTopic(name)

	return err
}

func InitAdmin() {

	adminConfig := sarama.NewConfig()

	host := os.Getenv("KAFKA_HOST")

	adminS, err := sarama.NewClusterAdmin([]string{host}, adminConfig)

	if err != nil {
		log.Fatalf("Erro on initialize admin %s", err.Error())
	}

	a.AdminImpl = adminS

}

func GetAdmin() AdminInterface {
	return &a
}
