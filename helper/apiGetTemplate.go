package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ApiGetRequest(url string, responseObject interface{}) {
	msgString := fmt.Sprintf("Performing GET %s ...", url)
	WriteStringToConsole(msgString)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	WriteStringToConsole("Decoding response JSON...")
	json.NewDecoder(resp.Body).Decode(responseObject)
}
