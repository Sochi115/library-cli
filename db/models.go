package db

type BookEntry struct {
	Key           string `db:"key"`
	Title         string `db:"title"`
	Authors       string `db:"authors"`
	Isbn13        string `db:"isbn_13"`
	Isbn10        string `db:"isbn_10"`
	Publishers    string `db:"publishers"`
	PublishDate   string `db:"publish_date"`
	NumberOfPages int    `db:"number_of_pages"`
	Rating        int    `db:"rating"`
}
