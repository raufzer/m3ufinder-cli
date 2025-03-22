package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Starts an interactive guide for the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		runInteractiveMode()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runInteractiveMode() {
	for {
		fmt.Println("Interactive mode started. Available commands:")
		fmt.Println("1. collect [category] - Extracts IPTV links from websites. Category is optional.")
		fmt.Println("2. save [format] - Exports extracted links in .m3u, .json, or .csv formats.")
		fmt.Println("3. exit - Exit interactive mode.")
		
		var command string
		fmt.Print("Enter command: ")
		fmt.Scanln(&command)
		
		switch command {
		case "1":
			fmt.Print("Enter category (optional): ")
			var category string
			fmt.Scanln(&category)
			if category != "" {
				collectCmd.Run(collectCmd, []string{category})
			} else {
				collectCmd.Run(collectCmd, []string{})
			}
		case "2":
			fmt.Print("Enter format (m3u, json, csv): ")
			var format string
			fmt.Scanln(&format)
			saveCmd.Run(saveCmd, []string{format})
		case "3":
			fmt.Println("Exiting interactive mode.")
			return
		default:
			fmt.Println("Invalid command. Please try again.")
		}
	}
}