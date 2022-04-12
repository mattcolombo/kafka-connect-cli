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
	Use:   "set",
	Short: "sets the log level set for a logger",
	Long:  "Allows to set the level set for a specific logger or connector plugin; allowed log levels are Java ones",
	Run: func(cmd *cobra.Command, args []string) {
		for _, host := range utilities.ConnectConfiguration.Hostnames {
			var loggerListURL string = host + "/admin/loggers/" + setPluginClass
			fmt.Println("--- Setting Log Level for Connect worker at", host, "---")
			fmt.Println("making a call to", loggerListURL) // control statement print - TOREMOVE
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
	LoggerSetCmd.Flags().StringVarP(&setPluginClass, "plugin-class", "", "", "plugin class to set the log level for (required)")
	LoggerSetCmd.Flags().StringVarP(&setLevel, "level", "", "", "log level to set (required)")
	LoggerSetCmd.MarkFlagRequired("plugin-class")
	LoggerSetCmd.MarkFlagRequired("level")
}

func buildSetPayload() []byte {
	payload, err := json.Marshal(map[string]interface{}{"level": setLevel})
	if err != nil {
		fmt.Printf("JSON build failed with error %s\n", err)
	}
	return payload
}
