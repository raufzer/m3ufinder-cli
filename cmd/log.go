package cmd

import (
    "m3ufinder-cli/internal/logger"

    "github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
    Use:   "log [level]",
    Short: "Displays logs (info, debug, error)",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        level := args[0]
        logger.DisplayLogs(level)
    },
}

func init() {
    rootCmd.AddCommand(logCmd)
}