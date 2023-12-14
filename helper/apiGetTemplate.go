package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func ApiGetRequest(url string, responseObject interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(responseObject)
}
