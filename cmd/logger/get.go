package logger

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var getPluginClass string

var LoggerGetCmd = &cobra.Command{
	Use:   "get",
	Short: "shows the log level set for a logger",
	Long:  "Allows to check the level set for a specific logger or connector plugin",
	Run: func(cmd *cobra.Command, args []string) {
		for _, host := range utilities.ConnectConfiguration.Hostnames {
			var loggerListURL string = host + "/admin/loggers/" + getPluginClass
			fmt.Println("--- Getting Log Level for Connect worker at", host, "---")
			fmt.Println("making a call to", loggerListURL) // control statement print - TOREMOVE
			response, err := utilities.DoCallByHost(http.MethodGet, loggerListURL, nil)
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				utilities.PrintResponseJson(response)
			}
		}
	},
}

func init() {
	LoggerGetCmd.Flags().StringVarP(&getPluginClass, "plugin-class", "", "", "plugin class to check the log level for (required)")
	LoggerGetCmd.MarkFlagRequired("plugin-class")
}
