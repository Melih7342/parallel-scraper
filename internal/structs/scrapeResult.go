package structs

type ScrapeResult struct {
	URL            string   `json:"url"`
	Title          string   `json:"title"`
	Description    string   `json:"description,omitempty"`
	Keywords       string   `json:"keywords,omitempty"`
	DeadLinks      []string `json:"dead_links"`
	DeadLinksCount int      `json:"dead_links_count"`
	Error          string   `json:"error"`
}
