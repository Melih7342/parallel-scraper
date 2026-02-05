package structs

type ScrapeResult struct {
	URL   string `json:"url"`
	Title string `json:"title"`
	Error string `json:"error"`
}
