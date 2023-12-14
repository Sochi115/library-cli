package save

import (
	"strings"

	"github.com/Sochi115/library-cli/db"
	"github.com/Sochi115/library-cli/helper"
)

var baseApiUrl = "https://openlibrary.org"

var bookMetaData = BookData{}

func HandleSaveBookByIsbn(isbn string) db.BookEntry {
	getIsbnData(isbn)
	authorKeys := getBookDataByWorks(bookMetaData.Key)
	getAuthors(authorKeys)

	return convertBookDataToBookEntry()
}

func getIsbnData(isbn string) {
	helper.ApiGetRequest(baseApiUrl+"/isbn/"+isbn+".json",
		&bookMetaData)
}

func getBookDataByWorks(works string) []map[string]string {
	responseObject := WorksResponseObject{}

	helper.ApiGetRequest(baseApiUrl+works+".json", &responseObject)

	return responseObject.Authors
}

func getAuthors(authorKeys []map[string]string) {
	var authorList []string
	for _, a := range authorKeys {
		responseObject := AuthorsResponseObject{}
		helper.ApiGetRequest(baseApiUrl+a["key"]+".json", &responseObject)

		authorList = append(authorList, responseObject.Name)
	}

	bookMetaData.Authors = authorList
}

func convertBookDataToBookEntry() db.BookEntry {
	return db.BookEntry{
		Key:           bookMetaData.Key,
		Title:         bookMetaData.Title,
		Authors:       strings.Join(bookMetaData.Authors, ", "),
		Isbn13:        strings.Join(bookMetaData.Isbn13, ", "),
		Isbn10:        strings.Join(bookMetaData.Isbn10, ", "),
		Publishers:    strings.Join(bookMetaData.Publishers, ", "),
		PublishDate:   bookMetaData.PublishDate,
		NumberOfPages: bookMetaData.NumberOfPages,
		Rating:        bookMetaData.Rating,
	}
}
