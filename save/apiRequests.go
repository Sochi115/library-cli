package save

import (
	"fmt"
	"strings"

	"github.com/Sochi115/library-cli/db"
	"github.com/Sochi115/library-cli/helper"
)

var baseApiUrl = "https://openlibrary.org"

var bookMetaData = BookData{}

func HandleSaveBookByIsbn(isbn string) db.BookEntry {
	key := getWorksIdFromIsbn(isbn)
	getBookDataByWorks(key)
	getAuthors()

	fmt.Println(bookMetaData)
	return convertBookDataToBookEntry()
}

func HandleSaveBookByWorks(works string) db.BookEntry {
	worksId := "/books/" + works

	getBookDataByWorks(worksId)
	getAuthors()
	fmt.Println(bookMetaData)
	return convertBookDataToBookEntry()
}

func getWorksIdFromIsbn(isbn string) string {
	var data map[string]string

	helper.ApiGetRequest(baseApiUrl+"/isbn/"+isbn+".json",
		&data)

	return data["key"]
}

func getBookDataByWorks(works string) {
	helper.ApiGetRequest(baseApiUrl+works+".json", &bookMetaData)
}

func getAuthors() {
	var authorList []string
	for _, a := range bookMetaData.AuthorKeys {

		var data map[string]string
		helper.ApiGetRequest(baseApiUrl+a.Key+".json", &data)

		authorList = append(authorList, data["name"])
	}

	bookMetaData.AuthorNames = authorList
}

func convertBookDataToBookEntry() db.BookEntry {
	return db.BookEntry{
		Key:           bookMetaData.Key,
		Title:         bookMetaData.Title,
		Authors:       strings.Join(bookMetaData.AuthorNames, ", "),
		Isbn13:        strings.Join(bookMetaData.Isbn13, ", "),
		Isbn10:        strings.Join(bookMetaData.Isbn10, ", "),
		Publishers:    strings.Join(bookMetaData.Publishers, ", "),
		PublishDate:   bookMetaData.PublishDate,
		NumberOfPages: bookMetaData.NumberOfPages,
		Rating:        bookMetaData.Rating,
	}
}
