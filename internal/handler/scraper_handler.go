package handler

import (
	"net/http"
	"time"

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

	start := time.Now()
	results := engine.RunEngine(input.URLs, input.Workers, input.Profile)
	duration := time.Since(start)

	c.JSON(200, gin.H{
		"total_results":  len(results),
		"execution_time": duration.String(),
		"data":           results,
	})
}
