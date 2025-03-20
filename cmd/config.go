package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
    Use:   "config [key] [value]",
    Short: "Sets global settings (parallel requests, rate-limiting, proxy)",
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        key := args[0]
        value := args[1]
        setConfig(key, value)
    },
}

func init() {
    rootCmd.AddCommand(configCmd)
}

func setConfig(key, value string) {
    fmt.Printf("Setting config %s to %s\n", key, value)
    // Implement config setting logic here
}