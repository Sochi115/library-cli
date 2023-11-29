package info

type BookInfo struct {
	Title       string   `json:"title"`
	PublishYear []int    `json:"publish_year"`
	Isbn        []string `json:"isbn"`
	Authors     []string `json:"author_name"`
}

type ApiSearchResponse struct {
	ResultCount int        `json:"num_found"`
	Docs        []BookInfo `json:"docs"`
}
