package list

import (
	"fmt"
	"strings"

	"github.com/alexeyco/simpletable"

	"github.com/Sochi115/library-cli/db"
	"github.com/Sochi115/library-cli/helper"
)

func DisplayAllSavedBooks(bookEntries []db.BookEntry) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: "Author(s)"},
			{Align: simpletable.AlignCenter, Text: "ISBN"},
		},
	}

	for i, b := range bookEntries {

		row := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", i+1)},
			{Align: simpletable.AlignLeft, Text: b.Title},
			{Align: simpletable.AlignLeft, Text: b.Authors},
			{Align: simpletable.AlignLeft, Text: generateIsbnString(b)},
		}
		table.Body.Cells = append(table.Body.Cells, row)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{
				Align: simpletable.AlignRight,
				Span:  4,
				Text:  fmt.Sprintf("Number of books: %d", len(bookEntries)),
			},
		},
	}

	helper.WriteStringToConsole(table.String())
}

func generateIsbnString(book db.BookEntry) string {
	isbns := []string{}

	if book.Isbn10 != "" {
		isbns = append(isbns, book.Isbn10)
	}
	if book.Isbn13 != "" {
		isbns = append(isbns, book.Isbn13)
	}
	return strings.Join(isbns, ", ")
}
