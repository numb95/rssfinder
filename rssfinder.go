package rssfinder

import (
	"golang.org/x/net/html"
	"net/http"
)

// FindRSSFeeds finds the RSS feeds in a webpage.
func FindRSSFeeds(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeeds []string
	var traverseHTML func(*html.Node)
	traverseHTML = func(n *html.Node) {
		// Check if the node is an <link> element
		if n.Type == html.ElementNode && n.Data == "link" {
			var feedURL string
			isRSS := false

			// Check the attributes of the <link> element
			for _, attr := range n.Attr {
				// Check if the type attribute is RSS or Atom
				if attr.Key == "type" && (attr.Val == "application/rss+xml" || attr.Val == "application/atom+xml") {
					isRSS = true
				} else if attr.Key == "href" {
					// Store the URL of the RSS feed
					feedURL = attr.Val
				}
			}

			// If it's an RSS or Atom feed and URL is not empty, add it to the list
			if isRSS && feedURL != "" {
				rssFeeds = append(rssFeeds, feedURL)
			}
		}

		// Recursively traverse the child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverseHTML(c)
		}
	}

	// Start traversing the HTML document from the root node
	traverseHTML(doc)

	return rssFeeds, nil
}
