package cmd

import (
	"fmt"
	"m3ufinder-cli/internal/saver"

	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save [format]",
	Short: "Exports extracted links in .m3u, .json, or .csv formats",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		format := args[0]
		err := saver.SaveM3U(m3uLinks, format)
		if err != nil {
			fmt.Println("Error saving links:", err)
		} else {
			fmt.Println("Links saved successfully.")
		}
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}