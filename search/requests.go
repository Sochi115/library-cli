package search

import "github.com/Sochi115/library-cli/helper"

var (
	apiBaseUrl   string = "https://openlibrary.org/search.json?"
	titleApiUrl  string = apiBaseUrl + "title="
	authorApiUrl string = apiBaseUrl + "author="
	apiFields    string = "&limit=25&language=eng"
)

func GetBookDataByTitle(title string) {
	queryString := parseInputQuery(title)
	apiUrl := titleApiUrl + queryString + apiFields
	books := SearchResponseObject{}

	helper.ApiGetRequest(apiUrl, &books)

	printBookTable(books)
}

func GetAuthorData(author string) {
	queryString := parseInputQuery(author)
	apiUrl := authorApiUrl + queryString + apiFields
	books := SearchResponseObject{}

	helper.ApiGetRequest(apiUrl, &books)

	printBookTable(books)
}
