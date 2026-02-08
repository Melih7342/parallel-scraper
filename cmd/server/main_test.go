package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Melih7342/parallel-scraper/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.StaticFile("/", "./public/dashboard.html")
	r.Static("/static", "./public/static")

	scraperHandler := handler.NewScraperHandler()

	api := r.Group("/api")
	{
		api.POST("/scrape", scraperHandler.PostScrape)
	}

	return r
}

func TestScrapeIntegration(t *testing.T) {
	router := setupRouter()

	// Case 1: Title Only Profile
	t.Run("Profile Title Only", func(t *testing.T) {
		body := map[string]interface{}{
			"urls":    []string{"https://google.com"},
			"workers": 1,
			"profile": "title",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/api/scrape", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}
	})

	// Case 2: Dead Links Profile
	t.Run("Profile Dead Links", func(t *testing.T) {
		body := map[string]interface{}{
			"urls":    []string{"https://google.com"},
			"workers": 1,
			"profile": "dead-links",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/api/scrape", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}
	})

	// Case 3: Title only with blank input
	t.Run("Profile Dead Links invalid input", func(t *testing.T) {
		body := map[string]interface{}{
			"urls":    []string{"hello"},
			"workers": 1,
			"profile": "dead-links",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/api/scrape", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &response)

		data := response["data"].([]interface{})
		firstResult := data[0].(map[string]interface{})

		if firstResult["error"] == "" {
			t.Errorf("Expected an error message for URL 'hello', but got none")
		}
	})

	// Case 4: Dead Links Profile
	t.Run("SEO Data", func(t *testing.T) {
		body := map[string]interface{}{
			"urls":    []string{"https://google.com"},
			"workers": 1,
			"profile": "seo",
		}
		jsonBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", "/api/scrape", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", w.Code)
		}
	})
}
