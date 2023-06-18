package indexer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
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

	timeoutStr := os.Getenv("INDEXER_TIMEOUT")

	timeconvert, err := strconv.ParseFloat(timeoutStr, 64)
	if err != nil {
		fmt.Println("Failed to parse INDEXER_TIMEOUT:", err)
		return
	}

	timeout = time.Duration(timeconvert * float64(time.Second))
}

// Index document in elasticsearch
func Index(document model.VideoData) {
	videoInfo := document.VideoInfo

	errorsValidate := videoInfo.Validate()

	if len(errorsValidate) > 0 {
		var errorMessages []string
		for _, err := range errorsValidate {
			errorMessages = append(errorMessages, err.Error())
		}

		fmt.Printf("Validate error for ID: %d\nMessage: %s\n", document.ID, strings.Join(errorMessages, "\n"))
		return
	}

	body, err := json.Marshal(videoInfo)

	if err != nil {
		fmt.Printf("Parse index error for KEY: %s\nMessage: %s\n", document.Key, err.Error())
		return
	}

	req := esapi.IndexRequest{
		DocumentID: fmt.Sprint(document.ID),
		Index:      videoInfo.Repository,
		Body:       bytes.NewReader(body),
		Refresh:    "true",
		Timeout:    timeout,
	}

	res, err := req.Do(context.Background(), es)

	if err != nil {
		fmt.Printf("index error for key: %s\nMessage: %s\n", document.Key, err.Error())
		return
	}

	if res.IsError() {
		log.Printf("[%s] Error indexing document", res.Status())
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
func Delete(data model.DeleteVideo) {

	// Create delete request
	req := esapi.DeleteRequest{
		Index:      data.Repository,
		DocumentID: fmt.Sprint(data.ID),
		Timeout:    timeout,
	}

	res, err := req.Do(context.Background(), es)

	// Handle with response
	if err != nil {
		fmt.Printf("Delete error for ID: %d\nMessage: %s\n", data.ID, err.Error())
		return
	}

	if res.IsError() {
		log.Printf("[%s] Error delete document, ID %d", res.Status(), data.ID)

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
