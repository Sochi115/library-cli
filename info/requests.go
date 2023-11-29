package info

import (
	"fmt"
)

var (
	apiBase       string = "https://openlibrary.org/search.json?"
	titleApiBase  string = apiBase + "title="
	authorApiBase string = apiBase + "author="
	apiFields     string = "&limit=25&language=eng"
)

func GetBookDataByTitle(title string) {
	queryString := parseInputQuery(title)
	api := titleApiBase + queryString + apiFields
	books := ApiSearchResponse{}

	apiGetRequest(api, &books)

	fmt.Println(books.ResultCount)
	printBookTable(books)
}

func GetAuthorData(author string) {
	queryString := parseInputQuery(author)
	api := authorApiBase + queryString + apiFields
	books := ApiSearchResponse{}

	apiGetRequest(api, &books)

	fmt.Println(books.ResultCount)
	printBookTable(books)
}
