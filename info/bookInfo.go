package info

import (
	"encoding/json"
	"fmt"
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

	books := OpenLibraryApiResponse{}

	fmt.Println("Decoding json")
	json.NewDecoder(res.Body).Decode(&books)

	fmt.Println(books.Docs[2])
}
