package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
    Use:   "run",
    Short: "Starts an interactive mode for users",
    Run: func(cmd *cobra.Command, args []string) {
        runInteractiveMode()
    },
}

func init() {
    rootCmd.AddCommand(runCmd)
}

func runInteractiveMode() {
    fmt.Println("Interactive mode started. Use commands like 'scrape', 'search', 'validate', 'save', etc.")
    // Implement interactive CLI logic here
}