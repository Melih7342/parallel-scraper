package scraper

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/Melih7342/parallel-scraper/internal/structs"
	"github.com/PuerkitoBio/goquery"
)

func ScrapePage(client *http.Client, urlStr string, profile string) structs.ScrapeResult {
	result := structs.ScrapeResult{
		URL:       urlStr,
		DeadLinks: []string{},
	}

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		result.Error = fmt.Sprintf("invalid URL: %v", err)
		return result
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/120.0.0.0")

	resp, err := client.Do(req)
	if err != nil {
		result.Error = fmt.Sprintf("network error: %v", err)
		return result
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		result.Error = fmt.Sprintf("HTTP error: %d %s", resp.StatusCode, resp.Status)
		return result
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		result.Error = fmt.Sprintf("parsing error: %v", err)
		return result
	}

	switch profile {
	case "seo":
		extractTitle(doc, &result)
		fillSEOData(doc, &result)
	case "dead-links":
		extractTitle(doc, &result)
		extractDeadLinks(doc, &result, client, urlStr)
	default:
		extractTitle(doc, &result)
	}
	return result
}

func extractTitle(doc *goquery.Document, res *structs.ScrapeResult) {
	res.Title = doc.Find("title").Text()
	if res.Title == "" {
		res.Title = doc.Find("h1").First().Text()
	}
	res.Title = strings.TrimSpace(res.Title)
}

func extractDeadLinks(doc *goquery.Document, res *structs.ScrapeResult, client *http.Client, baseURL string) {
	base, _ := url.Parse(baseURL)
	links := doc.Find("a")

	var wg sync.WaitGroup
	var mu sync.Mutex

	links.Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists || href == "" || href == "#" || strings.HasPrefix(href, "mailto:") || strings.HasPrefix(href, "tel:") {
			return
		}

		u, err := url.Parse(href)
		if err != nil {
			return
		}
		absoluteURL := base.ResolveReference(u).String()

		wg.Add(1)
		go func(link string) {
			defer wg.Done()

			req, err := http.NewRequest("HEAD", link, nil)
			if err != nil {
				return
			}
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/120.0.0.0")

			resp, err := client.Do(req)

			isDead := false
			if err != nil {
				isDead = true
			} else {
				defer resp.Body.Close()
				if resp.StatusCode >= 400 && resp.StatusCode != 405 {
					isDead = true
				}
			}

			if isDead {
				mu.Lock()
				res.DeadLinks = append(res.DeadLinks, link)
				res.DeadLinksCount++
				mu.Unlock()
			}
		}(absoluteURL)
	})

	wg.Wait()
}

func fillSEOData(doc *goquery.Document, res *structs.ScrapeResult) {
	res.Description, _ = doc.Find("meta[name='description']").Attr("content")
	res.Keywords, _ = doc.Find("meta[name='keywords']").Attr("content")
}
