package searcher

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/leocrispindev/streaming-video/searcher/internal/model"
	"github.com/leocrispindev/streaming-video/searcher/internal/util"
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

func Search(queryEngine model.Query) ([]model.Video, int, error) {

	var r map[string]interface{}

	var body bytes.Buffer

	// large scale  can be extremely dangerous and slow
	matchQuery := "match_all"

	if len(queryEngine.Searchers) > 0 {
		matchQuery = "match"
	}

	// Pagination
	queryEngine.From = (queryEngine.Page - 1) * queryEngine.Size

	query := map[string]interface{}{
		"size": queryEngine.Size,
		"from": queryEngine.From,
		"query": map[string]interface{}{
			matchQuery: queryEngine.Searchers,
		},
	}

	bodyQuery, err := json.Marshal(query)
	if err != nil {
		log.Fatalf("Error converting query to JSON: %s", err)
		return nil, 0, err
	}

	_, err = body.Write(bodyQuery)
	if err != nil {
		log.Fatalf("Error writing body: %s", err)
		return nil, 0, err
	}

	req := esapi.SearchRequest{
		Index:          []string{"video"},
		Body:           &body,
		Pretty:         true,
		SourceIncludes: queryEngine.Fields,
	}

	res, err := req.Do(context.Background(), es)

	//TODO tratar err
	if err != nil {
		return nil, 0, err
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

			return nil, 0, errors.New(msg)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	if err != nil {
		return nil, 0, err
	}

	// Exibir a string JSON resultante

	videosMap := []model.Video{}

	hits := r["hits"].(map[string]interface{})

	totalInterface := hits["total"].(map[string]interface{})["value"]

	total := util.ConvertToInt(totalInterface)

	println(total)

	for _, hit := range hits["hits"].([]interface{}) {
		source := hit.(map[string]interface{})["_source"]

		src, err := json.Marshal(source)
		if err != nil {
			fmt.Println("Erro ao converter para JSON:", err)
		}

		video := model.Video{}

		json.Unmarshal(src, &video)
		// Exibir a string JSON resultante
		videosMap = append(videosMap, video)
	}

	return videosMap, queryEngine.Page + 1, nil
}
