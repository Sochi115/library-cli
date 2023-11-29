package info

import (
	"fmt"
	"strings"

	"github.com/alexeyco/simpletable"
)

var titleApiBase string = "https://openlibrary.org/search.json?title="

func GetBookByTitle(title string) {
	queryString := parseInputQuery(title)
	api := titleApiBase + queryString + "&limit=25&language=eng"

	books := OpenLibraryApiResponse{}

	apiGetRequest(api, &books)

	fmt.Println(books.ResultCount)
	printTable(books)
}

func printTable(booksData OpenLibraryApiResponse) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: "Author(s)"},
			{Align: simpletable.AlignCenter, Text: "ISBN"},
		},
	}

	for _, b := range booksData.Docs {
		if len(b.Isbn) > 2 {
			b.Isbn = b.Isbn[:2]
		}
		row := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: b.Title},
			{Align: simpletable.AlignRight, Text: strings.Join(b.Authors, ", ")},
			{Align: simpletable.AlignRight, Text: strings.Join(b.Isbn, ", ")},
		}
		table.Body.Cells = append(table.Body.Cells, row)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{},
			{Align: simpletable.AlignRight, Text: "Result Count"},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", booksData.ResultCount)},
		},
	}

	table.SetStyle(simpletable.StyleCompactLite)

	writeToConsole(table.String())
}
