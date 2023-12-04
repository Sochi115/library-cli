package search

type BookInfo struct {
	Title       string   `json:"title"`
	PublishYear []int    `json:"publish_year"`
	Isbn        []string `json:"isbn"`
	Authors     []string `json:"author_name"`
}

type SearchResponseObject struct {
	ResultCount int        `json:"num_found"`
	Docs        []BookInfo `json:"docs"`
}
