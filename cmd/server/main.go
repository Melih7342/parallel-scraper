package main

import (
	"github.com/Melih7342/parallel-scraper/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.StaticFile("/", "./public/dashboard.html")
	r.Static("/static", "./public/static")

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
