package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	gUuid "github.com/google/uuid"
	"github.com/leocrispindev/streaming-video/searcher/internal/model"
	"github.com/leocrispindev/streaming-video/searcher/internal/service/searcher"
)

func Init() {
	searcher.Init()
}

func Search(resp http.ResponseWriter, req *http.Request) {

	uuid := gUuid.New().String()

	queryString := req.URL.Query()

	fields := []string{"title,category"}

	searchers := map[string]interface{}{}

	size := 10

	from := 0

	if queryString.Has("fields") {
		fields = strings.Split(queryString.Get("fields"), ",")
	}

	if queryString.Has("title") {
		searchers["title"] = queryString.Get("title")
	}

	if queryString.Has("synopsis") {
		searchers["synopsis"] = queryString.Get("synopsis")

	}

	if queryString.Has("size") {
		size, _ = strconv.Atoi(queryString.Get("size"))

	}

	if queryString.Has("from") {
		from, _ = strconv.Atoi(queryString.Get("from"))
		from += 1

	}

	searchQuery := model.Query{
		Fields:    fields,
		Searchers: searchers,
		Size:      size,
		From:      from,
	}

	docs, next, err := searcher.Search(searchQuery)

	resp.Header().Set("Content-type", "application/json")

	response := model.Response{
		Uuid: uuid,
	}

	if err != nil {
		response.Message = err.Error()

		body, _ := json.Marshal(response)

		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write(body)
		return
	} else {
		response.Message = "success"
		response.Docs = docs
		response.From = from

		body, _ := json.Marshal(response)

		resp.WriteHeader(http.StatusOK)

		resp.Write(body)
	}

}
