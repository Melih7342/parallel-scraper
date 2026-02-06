package main

import (
	"github.com/Melih7342/parallel-scraper/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	scraperHandler := handler.NewScraperHandler()

	api := r.Group("/api")
	{
		api.POST("/scrape/title", scraperHandler.PostScrape)
	}

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
