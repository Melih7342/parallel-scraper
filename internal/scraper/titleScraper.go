package scraper

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func titleScraper(url string) (string, error) {
	// 1. Request with timeout
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// 2. Load html doc
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	// 3. Find the title (<title> or <h1>)
	title := doc.Find("title").Text()
	if title == "" {
		title = doc.Find("h1").First().Text()
	}

	return strings.TrimSpace(title), nil
}
