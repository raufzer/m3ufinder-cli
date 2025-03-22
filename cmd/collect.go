package cmd

import (
	"fmt"
	"m3ufinder-cli/internal/scraper"

	"github.com/spf13/cobra"
)

var collectCmd = &cobra.Command{
	Use:   "collect [category]",
	Short: "Extracts IPTV links from websites",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var urls []string
		if len(args) == 1 {
			category := args[0]
			url := scraper.GetCategoryURL(category)
			if url != "" {
				urls = append(urls, url)
			} else {
				fmt.Println("Invalid category. Available categories: sports, news, entertainment")
				return
			}
		} else {
			urls = scraper.GetAllURLs()
		}

		var allLinks []string
		for _, url := range urls {
			links, err := scraper.ScrapeIPTVLinks(url)
			if err != nil {
				fmt.Println("Error scraping URL:", err)
				continue
			}
			allLinks = append(allLinks, links...)
		}

		for _, link := range allLinks {
			fmt.Println("Found IPTV link:", link)
		}

		// Save links to a global variable for later use
		m3uLinks = allLinks

		// Indicate that scraping is done
		fmt.Println("Scraping completed. Use the save command to save the links.")
	},
}

var m3uLinks []string

func init() {
	scraper.LoadEnv()
	rootCmd.AddCommand(collectCmd)
}