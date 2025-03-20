package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"m3ufinder-cli/internal/scraper"
)

// scrapeCmd represents the scrape command
var scrapeCmd = &cobra.Command{
	Use:   "scrape [url]",
	Short: "Extracts M3U links from a website",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		links, err := scraper.ScrapeM3U(url)
		if err != nil {
			fmt.Println("Error scraping URL:", err)
			return
		}
		for _, link := range links {
			fmt.Println("Found M3U link:", link)
		}
	},
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
}
