package save

import "github.com/Sochi115/library-cli/helper"

var isbnBaseUrl = "https://openlibrary.org/isbn/"

func GetIsbnData(isbn string) IsbnResponseObject {
	responseObject := IsbnResponseObject{}

	helper.ApiGetRequest(generateIsbnUrl(isbn), &responseObject)

	return responseObject
}

func generateIsbnUrl(isbn string) string {
	return isbnBaseUrl + isbn + ".json"
}
