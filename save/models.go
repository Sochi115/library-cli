package save

type IsbnResponseObject struct {
	Key           string   `json:"key"`
	Title         string   `json:"title"`
	NumberOfPages int      `json:"number_of_pages"`
	Isbn13        []string `json:"isbn_13"`
	Isbn10        []string `json:"isbn_10"`
}
