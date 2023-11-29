package info

type BookInfo struct {
	Title         string   `json:"title"`
	PublishYear   []int    `json:"publish_year"`
	NumberOfPages int      `json:"number_of_pages_median"`
	Isbn          []string `json:"isbn"`
	Authors       []string `json:"author_name"`
}

type OpenLibraryApiResponse struct {
	ResultCount int        `json:"num_found"`
	Docs        []BookInfo `json:"docs"`
}
