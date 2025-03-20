package scraper

import (
	"os"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
	"strings"
)

// LoadEnv loads the environment variables from the .env file.
func LoadEnv() error {
	return godotenv.Load()
}

// GetCategoryURL returns the URL for a given category.
func GetCategoryURL(category string) string {
	switch strings.ToLower(category) {
	case "sports":
		return os.Getenv("SPORTS_URL")
	case "news":
		return os.Getenv("NEWS_URL")
	case "entertainment":
		return os.Getenv("ENTERTAINMENT_URL")
	default:
		return ""
	}
}

// GetAllURLs returns all URLs from the environment variables.
func GetAllURLs() []string {
	return []string{
		os.Getenv("WATCH_IPTV"),
		os.Getenv("JACKAL"),
		os.Getenv("IPTV_PLAYER"),
		os.Getenv("WORLDS_TV_MOBILE"),
		os.Getenv("TV_TUNER"),
		os.Getenv("EBK_IPTV"),
		os.Getenv("TV_GARDEN"),
		os.Getenv("GLOBAL_FREE_TV"),
		os.Getenv("LYNGSAT_STREAM"),
		os.Getenv("FREETUXTV_WEBTV_MANAGER"),
		os.Getenv("CXTV"),
		os.Getenv("COSTA_RICA_EN_VIVO"),
		os.Getenv("DOMINICANA_EN_VIVO"),
		os.Getenv("HONDURAS_EN_VIVO"),
		os.Getenv("GUATEMALA_EN_VIVO"),
		os.Getenv("EL_SALVADOR"),
		os.Getenv("BOLIVIA_EN_VIVO"),
		os.Getenv("HAITI_BROADCASTING"),
		os.Getenv("SQUID_TV"),
		os.Getenv("ONLINE_TV"),
		os.Getenv("TDTCHANNELS"),
		os.Getenv("ONLINESTREAM_LIVE"),
		os.Getenv("EASY_WEB_TV"),
		os.Getenv("TNT_EN_DIRECT"),
		os.Getenv("EPG_PW"),
		os.Getenv("CHINESE_TV"),
		os.Getenv("PHOTOCALL_TV"),
		os.Getenv("M3UPT"),
	}
}

// ScrapeIPTVLinks scrapes M3U links from the provided URL.
func ScrapeIPTVLinks(url string) ([]string, error) {
	var links []string
	c := colly.NewCollector()

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

// isM3U checks if a link is an M3U or M3U8 playlist.
func isM3U(link string) bool {
	link = strings.ToLower(link)
	return strings.HasSuffix(link, ".m3u") || strings.HasSuffix(link, ".m3u8")
}