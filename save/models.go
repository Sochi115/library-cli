package save

type BookData struct {
	Key           string                  `json:"key"`
	Title         string                  `json:"title"`
	AuthorKeys    []AuthorsResponseObject `json:"authors"`
	AuthorNames   []string                `json:"-"`
	Isbn13        []string                `json:"isbn_13"`
	Isbn10        []string                `json:"isbn_10"`
	Publishers    []string                `json:"-"`
	PublishDate   string                  `json:"-"`
	NumberOfPages int                     `json:"number_of_pages"`
	Rating        int                     `json:"-"`
}

type AuthorsResponseObject struct {
	Key string `json:"key"`
}
