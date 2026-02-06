package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Melih7342/parallel-scraper/internal/structs"
	"github.com/PuerkitoBio/goquery"
)

func ScrapePage(client *http.Client, url string) structs.ScrapeResult {
	// 1. Initialize result struct
	result := structs.ScrapeResult{URL: url}

	// 2. Request
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/120.0.0.0")
	resp, err := client.Do(req)
	if err != nil {
		result.Error = fmt.Sprintf("network error: %v", err)
		return result
	}

	defer resp.Body.Close()

	// 3. Check status code
	if resp.StatusCode != http.StatusOK {
		result.Error = fmt.Sprintf("HTTP error: %d %s", resp.StatusCode, resp.Status)
		return result
	}

	// 4. Load html document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		result.Error = fmt.Sprintf("parsing error: %v", err)
		return result
	}

	// 5. Find the title
	title := doc.Find("title").Text()
	if title == "" {
		title = doc.Find("h1").First().Text()
	}

	result.Title = strings.TrimSpace(title)
	return result
}
