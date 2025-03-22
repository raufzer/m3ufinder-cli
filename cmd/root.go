package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "m3ufinder",
	Short: "M3U Finder CLI",
	Long:  "M3U Finder CLI is a tool to extract and validate M3U links",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Define groups for commands
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(collectCmd)
	rootCmd.AddCommand(saveCmd)
}