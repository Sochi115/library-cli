package save

import (
	"fmt"

	"github.com/Sochi115/library-cli/helper"
)

var baseApiUrl = "https://openlibrary.org"

var bookMetaData = BookMetaData{}

func HandleSaveBookByIsbn(isbn string) {
	GetIsbnData(isbn)
	authorKeys := GetBookDataByWorks(bookMetaData.Key)
	GetAuthors(authorKeys)

	fmt.Println(bookMetaData)
}

func GetIsbnData(isbn string) {
	helper.ApiGetRequest(baseApiUrl+"/isbn/"+isbn+".json",
		&bookMetaData)
}

func GetBookDataByWorks(works string) []map[string]string {
	responseObject := WorksResponseObject{}

	helper.ApiGetRequest(baseApiUrl+works+".json", &responseObject)

	return responseObject.Authors
}

func GetAuthors(authorKeys []map[string]string) {
	var authorList []string
	for _, a := range authorKeys {
		responseObject := AuthorsResponseObject{}
		helper.ApiGetRequest(baseApiUrl+a["key"]+".json", &responseObject)

		authorList = append(authorList, responseObject.Name)
	}

	bookMetaData.Authors = authorList
}
