package handler

import (
	"net/http"

	"github.com/Melih7342/parallel-scraper/internal/engine"
	"github.com/Melih7342/parallel-scraper/internal/structs"
	"github.com/gin-gonic/gin"
)

type ScraperHandler struct{}

func NewScraperHandler() *ScraperHandler {
	return &ScraperHandler{}
}

func (h *ScraperHandler) PostScrape(c *gin.Context) {
	var input structs.ScrapeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	if input.Workers <= 0 {
		input.Workers = 5
	}

	results := engine.RunEngine(input.URLs, input.Workers)

	c.JSON(http.StatusOK, results)
}
