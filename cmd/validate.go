package cmd

import (
    "fmt"
    "m3ufinder-cli/internal/validator"

    "github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
    Use:   "validate",
    Short: "Checks if extracted M3U links are active",
    Run: func(cmd *cobra.Command, args []string) {
        validLinks := validator.ValidateM3U(m3uLinks)
        for _, link := range validLinks {
            fmt.Println("Valid:", link)
        }
    },
}

func init() {
    rootCmd.AddCommand(validateCmd)
}