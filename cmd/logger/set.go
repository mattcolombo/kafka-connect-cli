package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var setPluginClass, setLevel string

var LoggerSetCmd = &cobra.Command{
	Use:   "set [flags] logger_name",
	Short: "sets the log level set for a logger",
	Long:  "Allows to set the log level set for a specific logger or connector plugin at runtime; allowed log levels are as in Java [\"OFF\",\"FATAL\",\"ERROR\",\"WARN\",\"INFO\",\"DEBUG\",\"TRACE\"]",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		setPluginClass = args[0]
		for _, host := range utilities.ConnectConfiguration.Hostnames {
			var loggerListURL string = host + "/admin/loggers/" + setPluginClass
			fmt.Println("--- Setting Log Level", setLevel, "for Connect worker at", host, "---")
			//fmt.Println("making a call to", loggerListURL) // control statement print
			response, err := utilities.DoCallByHost(http.MethodPut, loggerListURL, bytes.NewBuffer(buildSetPayload()))
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				utilities.PrintResponseJson(response)
			}
		}
	},
}

func init() {
	//LoggerSetCmd.Flags().StringVarP(&setPluginClass, "plugin-class", "", "", "plugin class to set the log level for (required)")
	LoggerSetCmd.Flags().StringVarP(&setLevel, "level", "", "ERROR", "log level to set")
	//LoggerSetCmd.MarkFlagRequired("plugin-class")
	//LoggerSetCmd.MarkFlagRequired("level")
}

func buildSetPayload() []byte {
	payload, err := json.Marshal(map[string]interface{}{"level": setLevel})
	if err != nil {
		fmt.Printf("JSON build failed with error %s\n", err)
	}
	return payload
}
