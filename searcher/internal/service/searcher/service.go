package searcher

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/leocrispindev/streaming-video/searcher/internal/model"
)

var es *elastic.Client

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
}

func Search(queryEngine model.Query) ([]map[string]interface{}, error) {

	var r map[string]interface{}

	var body bytes.Buffer

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": queryEngine.Searchers,
		},
	}

	bodyQuery, err := json.Marshal(query)
	if err != nil {
		log.Fatalf("Error converting query to JSON: %s", err)
		return nil, err
	}

	_, err = body.Write(bodyQuery)
	if err != nil {
		log.Fatalf("Error writing body: %s", err)
		return nil, err
	}

	req := esapi.SearchRequest{
		Index:          []string{"video"},
		Body:           &body,
		Pretty:         true,
		SourceIncludes: queryEngine.Fields,
	}

	println(strings.Join(req.StoredFields, ","))

	res, err := req.Do(context.Background(), es)

	//TODO tratar err
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			msg := fmt.Sprintf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)

			return nil, errors.New(msg)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	if err != nil {
		return nil, err
	}

	jsonString, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
	}

	// Exibir a string JSON resultante
	fmt.Println(string(jsonString))

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		jsonString, err := json.Marshal(hit)
		if err != nil {
			fmt.Println("Erro ao converter para JSON:", err)
		}

		// Exibir a string JSON resultante
		fmt.Println(string(jsonString))
	}

	return []map[string]interface{}{}, nil
}
