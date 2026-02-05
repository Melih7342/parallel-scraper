package structs

type ScrapeInput struct {
	URL     []string `json:"url" binding:"required"`
	Workers int      `json:"workers"`
}
