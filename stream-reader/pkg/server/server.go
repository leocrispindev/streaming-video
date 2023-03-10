package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NygmaC/streamming-video/stream-reader/internal/model"
)

func Init() {

	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {

		var process model.Proccess

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("could not read body: %s\n", err)
		}

		json.Unmarshal(body, &process)

	})

}
