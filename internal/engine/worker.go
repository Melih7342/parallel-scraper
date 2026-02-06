package engine

import (
	"net/http"

	"github.com/Melih7342/parallel-scraper/internal/scraper"
	"github.com/Melih7342/parallel-scraper/internal/structs"
)

func StartWorker(id int, jobs <-chan string, results chan<- structs.ScrapeResult, client *http.Client, profile string) {
	for url := range jobs {
		result := scraper.ScrapePage(client, url, profile)
		results <- result
	}
}
