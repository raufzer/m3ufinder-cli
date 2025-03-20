package scraper

import (
	"github.com/gocolly/colly"
)

func ScrapeM3U(url string) ([]string, error) {
	c := colly.NewCollector()
	var links []string

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if isM3U(link) {
			links = append(links, link)
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
        return 
	})

	c.Visit(url)
	return links, nil
}

func isM3U(link string) bool {
	return len(link) > 4 && (link[len(link)-4:] == ".m3u" || link[len(link)-5:] == ".m3u8")
}
