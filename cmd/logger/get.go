package logger

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

//var getPluginClass string

var LoggerGetCmd = &cobra.Command{
	Use:   "get [flags] logger_name",
	Short: "shows the log level set for a logger",
	Long:  "Allows to check the log level set for a specific logger or connector plugin",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pluginClass = args[0]
		fmt.Printf("pluiginClass is %s", pluginClass) // control statement print
		for _, host := range utilities.ConnectConfiguration.Hostnames {
			var loggerListURL = fmt.Sprintf("%s/admin/loggers/%s", host, pluginClass)
			fmt.Printf("--- Getting Log Level for Connect worker at %s ---", host)
			//fmt.Println("making a call to", loggerListURL) // control statement print
			response, err := utilities.DoCallByHost(http.MethodGet, loggerListURL, nil)
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				utilities.PrintResponseJson(response)
			}
		}
	},
}
