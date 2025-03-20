package cmd

import (
    "fmt"
    "m3ufinder-cli/internal/searcher"

    "github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
    Use:   "search [keyword or glob pattern]",
    Short: "Filters M3U links based on keywords or glob pattern",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        pattern := args[0]
        filteredLinks := searcher.SearchM3U(m3uLinks, pattern)
        for _, link := range filteredLinks {
            fmt.Println(link)
        }
    },
}

func init() {
    rootCmd.AddCommand(searchCmd)
}

var m3uLinks []string // This should be populated with the scraped links