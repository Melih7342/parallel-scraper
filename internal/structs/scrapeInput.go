package structs

type ScrapeInput struct {
	URL     string `json:"url"`
	Workers int    `json:"workers"`
}
