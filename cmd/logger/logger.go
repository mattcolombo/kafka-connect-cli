package logger

import (
	"github.com/spf13/cobra"
)

var LoggerCmd = &cobra.Command{
	Use:   "logger",
	Short: "logger short description",
	Long:  "logger long description",
}

func init() {
	LoggerCmd.AddCommand(LoggerListCmd)
	LoggerCmd.AddCommand(LoggerGetCmd)
	LoggerCmd.AddCommand(LoggerSetCmd)
}
