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
	traverseHTML(doc, &rssFeeds)
	return rssFeeds, nil
}

func traverseHTML(n *html.Node, rssFeeds *[]string) {
	if n.Type == html.ElementNode && n.Data == "link" {
		feedURL := getFeedURL(n)
		if feedURL != "" {
			*rssFeeds = append(*rssFeeds, feedURL)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseHTML(c, rssFeeds)
	}
}

func getFeedURL(n *html.Node) string {
	var feedURL string
	isRSS := false
	for _, attr := range n.Attr {
		if attr.Key == "type" && (attr.Val == "application/rss+xml" || attr.Val == "application/atom+xml") {
			isRSS = true
		} else if attr.Key == "href" {
			feedURL = attr.Val
		}
	}
	if isRSS {
		return feedURL
	}
	return ""
}
