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
			"http://localhost:9200",
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

func Index(document model.Video) {

	errorsValidate := document.Validate()

	if len(errorsValidate) > 0 {
		var errorMessages []string
		for _, err := range errorsValidate {
			errorMessages = append(errorMessages, err.Error())
		}

		fmt.Printf("Validate error for ID: %d\nMessage: %s\n", document.ID, strings.Join(errorMessages, "\n"))
		return
	}

	body, err := json.Marshal(document)

	if err != nil {
		fmt.Printf("Parse index error for ID: %d\nMessage: %s\n", document.ID, err.Error())
		return
	}

	// Set up the request object.
	req := esapi.IndexRequest{
		DocumentID: fmt.Sprint(document.ID),
		Index:      document.Repository,
		Body:       bytes.NewReader(body),
		Refresh:    "true",
		Timeout:    timeout,
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), es)

	if err != nil {
		fmt.Printf("index error for ID: %d\nMessage: %s\n", document.ID, err.Error())
		return
	}

	if res.IsError() {
		log.Printf("[%s] Error indexing document", res.Status())
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
}
