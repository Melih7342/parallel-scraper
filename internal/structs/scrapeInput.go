package structs

type ScrapeInput struct {
	URLs    []string `json:"urls" binding:"required"`
	Workers int      `json:"workers"`
}
