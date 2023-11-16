package info

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

var titleApiBase string = "https://openlibrary.org/search.json?title="

func GetBookByTitle(title string) {
	tokens := strings.Split(strings.TrimSpace(title), " ")

	queryString := tokens[0]

	for _, token := range tokens[1:] {
		queryString += "+"
		queryString += token
	}

	api := titleApiBase + queryString
	fmt.Println("Making request to", api)
	res, err := http.Get(titleApiBase + queryString)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
