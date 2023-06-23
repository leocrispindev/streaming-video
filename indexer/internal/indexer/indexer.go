package indexer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/leocrispindev/streaming-video/indexer/internal/model"
)

var es *elastic.Client
var timeout time.Duration

func Init() {
	cfg := elastic.Config{
		Addresses: []string{
			os.Getenv("ELASTICSEARCH_HOST"),
		},
		// ...
	}

	client, err := elastic.NewClient(cfg)

	if err != nil {
		log.Fatal(err)
	}

	es = client

	// timeoutStr := os.Getenv("INDEXER_TIMEOUT")

	// timeconvert, err := strconv.ParseInt(timeoutStr, 0, 32)
	// if err != nil {
	// 	fmt.Println("Failed to parse INDEXER_TIMEOUT:", err)
	// 	return
	// }
	// timeout = time.Duration(timeconvert) * time.Second
}

// Index document in elasticsearch
func Index(document model.Document) {

	errorsValidate := document.Validate()

	if len(errorsValidate) > 0 {
		var errorMessages []string
		for _, err := range errorsValidate {
			errorMessages = append(errorMessages, err.Error())
		}

		fmt.Printf("Validate error for KEY: %s\nMessage: %s\n", document.Key, strings.Join(errorMessages, "\n"))
		return
	}

	videoInfo := document.VideoInfo

	errorsValidate = videoInfo.Validate()

	if len(errorsValidate) > 0 {
		var errorMessages []string
		for _, err := range errorsValidate {
			errorMessages = append(errorMessages, err.Error())
		}

		fmt.Printf("Validate error for KEY: %s\nMessage: %s\n", document.Key, strings.Join(errorMessages, "\n"))
		return
	}

	body, err := json.Marshal(videoInfo)

	if err != nil {
		fmt.Printf("Parse index error for KEY: %s\nMessage: %s\n", document.Key, err.Error())
		return
	}

	println(timeout.String())

	req := esapi.IndexRequest{
		DocumentID: document.Key,
		Index:      document.Repository,
		Body:       bytes.NewReader(body),
		Refresh:    "true",
		//Timeout:    timeout.String(),
	}

	res, err := req.Do(context.Background(), es)

	if err != nil {
		fmt.Printf("index error for key: %s\nMessage: %s\n", document.Key, err.Error())
		return
	}

	if res.IsError() {
		log.Printf("[%s] Error indexing document %s", res.Status(), res.String())
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)

		} else {
			log.Printf("index success for [KEY]=%s [%s] %s; version=%d", document.Key, res.Status(), r["result"], int(r["_version"].(float64)))

		}
	}
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
}

// Delete document from elasticsearch
func Delete(document model.Document) {

	errorsValidate := document.Validate()

	if len(errorsValidate) > 0 {
		var errorMessages []string
		for _, err := range errorsValidate {
			errorMessages = append(errorMessages, err.Error())
		}

		fmt.Printf("Delete Validate error for KEY: %s\nMessage: %s\n", document.Key, strings.Join(errorMessages, "\n"))
		return
	}
	// Create delete request
	req := esapi.DeleteRequest{
		Index:      document.Repository,
		DocumentID: document.Key,
	}

	res, err := req.Do(context.Background(), es)

	// Handle with response
	if err != nil {
		fmt.Printf("Delete error for KEY: %s\nMessage: %s\n", document.Key, err.Error())
		return
	}

	if res.IsError() {
		log.Printf("[%s] Error delete document, KEY %s", res.Status(), document.Key)

	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)

		} else {
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))

		}
	}
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

}
