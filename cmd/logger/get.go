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
	Short: "logger get short description",
	Long:  "logger get long description",
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
	LoggerGetCmd.Flags().StringVarP(&getPluginClass, "plugin-class", "", "", "plugin class to check for log level (required)")
	LoggerGetCmd.MarkFlagRequired("plugin-class")
}
