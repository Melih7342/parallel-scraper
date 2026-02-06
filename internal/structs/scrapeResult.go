package structs

type ScrapeResult struct {
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Keywords    string `json:"keywords,omitempty"`
	Server      string `json:"server,omitempty"`
	Error       string `json:"error"`
}
