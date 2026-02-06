package engine

import (
	"net/http"
	"sync"
	"time"

	"github.com/Melih7342/parallel-scraper/internal/structs"
)

func RunEngine(urls []string, workerCount int, profile string) []structs.ScrapeResult {
	jobs := make(chan string, len(urls))
	results := make(chan structs.ScrapeResult, len(urls))

	var wg sync.WaitGroup
	client := &http.Client{Timeout: time.Second * 10}

	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			StartWorker(workerID, jobs, results, client, profile)
		}(w)
	}
	for _, url := range urls {
		jobs <- url
	}
	close(jobs)
	wg.Wait()
	close(results)

	finalResults := make([]structs.ScrapeResult, 0, len(urls))
	for result := range results {
		finalResults = append(finalResults, result)
	}
	return finalResults
}
