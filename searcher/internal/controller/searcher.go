package controller

import (
	"net/http"
)

func Search(resp http.ResponseWriter, req *http.Request) {

	queryString := req.URL.Query()
	println(queryString)

	//searcher.Search()

}
