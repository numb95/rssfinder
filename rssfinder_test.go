package rssfinder

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRSSFeeds(t *testing.T) {
	// Create a test server to mock the HTTP response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		// Sample HTML response containing RSS feed links
		htmlResponse := `
			<html>
			<head>
				<link rel="alternate" type="application/rss+xml" href="https://example.com/feed.xml">
				<link rel="alternate" type="application/rss+xml" href="https://example.com/rss.xml">
				<link rel="alternate" type="application/atom+xml" href="https://example.com/atom.xml">
				<link rel="stylesheet" href="styles.css">
			</head>
			<body>
				<h1>Test Page</h1>
			</body>
			</html>
		`
		w.Write([]byte(htmlResponse))
	}))
	defer server.Close()

	// Set the test server URL as the base URL for the test
	baseURL := server.URL

	// Call the function being tested
	rssFeeds, err := FindRSSFeeds(baseURL)
	assert.NoError(t, err, "Unexpected error")

	// Assert the expected RSS feed URLs
	expectedFeeds := []string{
		"https://example.com/feed.xml",
		"https://example.com/rss.xml",
		"https://example.com/atom.xml",
	}
	assert.Equal(t, expectedFeeds, rssFeeds, "Unexpected RSS feed URLs")
}
