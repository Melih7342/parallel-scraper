# 🚀 Go-Scraper Engine

A high-performance, parallel web scraping platform built with **Go** and **Gin**. This engine allows users to analyze multiple URLs simultaneously, extract SEO metadata, and identify dead links in real-time using Go's powerful concurrency primitives.


## ✨ Features

- **Concurrent Scraping:** Leveraging Goroutines and Channels to process massive URL lists with minimal latency.
- **Smart Profiling System:**
    - `Title Only`: Lightweight check for page availability.
    - `SEO Meta-Data`: Extraction of Meta-Titles, Descriptions, and Keywords.
    - `Dead Link Checker`: Deep-scan of all page anchors to detect 404/500 errors using a parallel verification pool.
- **Modern Dashboard:** A clean, responsive interface built with Tailwind CSS, featuring real-time statistics and status indicators.
- **Robust Engine Logic:** - Automatic resolution of relative paths to absolute URLs.
    - User-Agent spoofing to bypass basic anti-bot headers.
    - **Two-Step Verification:** Parallel `HEAD` requests with a `GET` fallback for accurate link validation.

## 🛠 Tech Stack

- **Backend:** [Go](https://golang.org/) (Golang)
- **Web Framework:** [Gin Gonic](https://gin-gonic.com/)
- **Scraping Library:** [Goquery](https://github.com/PuerkitoBio/goquery)
- **Frontend:** HTML5, Vanilla JavaScript (Async/Fetch), [Tailwind CSS](https://tailwindcss.com/)

## 🚀 Quick Start

### Prerequisites
- Go 1.21 or higher installed on your system.

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/Melih7342/parallel-scraper.git
   cd parallel-scraper

2. Install dependencies:
    ```bash
    go mod download

3. Start the server:

    ```bash
    go run cmd/server/main.go
   
4. Access the Dashboard: Open your browser and navigate to http://localhost:8080

## 🧪 Testing & Debugging

### To run the automated integration tests:

    go test -v ./...

## 📁 Project Structure
```Plaintext
├── cmd/
│   └── server/main.go          # Entry Point & Router Configuration
├── internal/
│   ├── engine/                 # Worker-Pool & Concurrency Management
│   ├── handler/                # API Request/Response Logic
│   ├── scraper/                # Core Scraping & Link Validation Logic
│   └── structs/                # Shared Data Models (JSON Mappings)
├── public/
│   ├── static/                 # JavaScript (dashboard.js)
│   └── dashboard.html          # Main Frontend UI
├── go.mod                      # Go Dependencies
└── README.md                   # Project Documentation
```

## 🧠Architecture Highlights
The engine uses a Worker Pool Pattern:

- **Jobs Channel:** Receives incoming URLs from the API.

- **Workers:** A configurable number of Goroutines process the jobs in parallel.

- **Results Channel:** Collects scraped data and pushes it back to the handler for the final JSON response.

- **Link Validation:** Inside each worker, links are checked concurrently using a sync.WaitGroup and sync.Mutex to maximize throughput without race conditions.
