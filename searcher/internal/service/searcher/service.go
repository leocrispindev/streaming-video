package searcher

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/leocrispindev/streaming-video/searcher/internal/model"
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
}

func Search(queryEngine model.Query) ([]map[string]interface{}, error) {

	var r map[string]interface{}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": queryEngine.Searchers,
		},
	}

	strQuery, err := json.Marshal(query)

	if err != nil {
		log.Fatalf("Error on convert query", err)
		return nil, err
	}

	req := esapi.SearchRequest{
		Index:        []string{"video"},
		StoredFields: queryEngine.Fields,
		Query:        string(strQuery),
	}

	res, err := req.Do(context.Background(), es)

	//TODO tratar err

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	//TODO tratar os HITS

	return nil, nil
}
