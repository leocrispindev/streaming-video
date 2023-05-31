package indexer

import (
	"log"

	elastic "github.com/elastic/go-elasticsearch/v8"
)

func Init() {
	cfg := elastic.Config{
		Addresses: []string{
			"https://localhost:9200",
			"https://localhost:9201",
		},
		// ...
	}

	client, err := elastic.NewClient(cfg)

	if err != nil {
		log.Fatal(err)
	}
}

func Index(topic string, body string) {

}
