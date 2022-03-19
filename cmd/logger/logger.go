package logger

import (
	"github.com/spf13/cobra"
)

var LoggerCmd = &cobra.Command{
	Use:   "logger",
	Short: "short description",
	Long:  "long description",
}

func init() {
	LoggerCmd.AddCommand(LoggerListCmd)
}
