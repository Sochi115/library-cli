package search

import (
	"fmt"
	"strings"

	"github.com/alexeyco/simpletable"

	"github.com/Sochi115/library-cli/helper"
)

func printBookTable(booksData SearchResponseObject) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Title"},
			{Align: simpletable.AlignCenter, Text: "Author(s)"},
			{Align: simpletable.AlignCenter, Text: "ISBN"},
			{Align: simpletable.AlignCenter, Text: "Years Published"},
		},
	}

	for i, b := range booksData.Docs {
		row := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", i+1)},
			{Align: simpletable.AlignRight, Text: longStringToMultiline(b.Title, 5)},
			{
				Align: simpletable.AlignRight,
				Text:  longStringToMultiline(strings.Join(b.Authors, ", "), 2),
			},
			{
				Align: simpletable.AlignRight,
				Text:  longStringToMultiline(strings.Join(b.Isbn, ", "), 2),
			},
			{
				Align: simpletable.AlignRight,
				Text:  longStringToMultiline(intSliceToString(b.PublishYear), 3),
			},
		}
		table.Body.Cells = append(table.Body.Cells, row)
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Span: 4, Text: "Result Count"},
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", booksData.ResultCount)},
		},
	}

	table.SetStyle(simpletable.StyleUnicode)

	helper.WriteStringToConsole(table.String())
}
