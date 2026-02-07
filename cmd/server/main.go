package main

import (
	"github.com/Melih7342/parallel-scraper/internal/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	scraperHandler := handler.NewScraperHandler()

	api := r.Group("/api")
	{
		api.POST("/scrape", scraperHandler.PostScrape)
	}

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
