package logger

import (
	"github.com/spf13/cobra"
)

var pluginClass string

var LoggerCmd = &cobra.Command{
	Use:   "logger",
	Short: "manage loggers at runtime",
	Long: `Allows to manage loggers and log levels at runtime without worker restart. 
---
Because loggers and log levels can be set independently for each worker, this command will make the same 
call to all the Connect workers listed in the hostnames section of the configuration file`,
}

func init() {
	LoggerCmd.AddCommand(LoggerListCmd)
	LoggerCmd.AddCommand(LoggerGetCmd)
	LoggerCmd.AddCommand(LoggerSetCmd)
}
