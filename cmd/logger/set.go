package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var setPluginClass, setLevel string

var LoggerSetCmd = &cobra.Command{
	Use:   "set",
	Short: "logger set short description",
	Long:  "logger set long description",
	Run: func(cmd *cobra.Command, args []string) {
		for _, host := range utilities.ConnectConfiguration.Hostname {
			var loggerListURL string = host + "/admin/loggers/" + setPluginClass
			fmt.Println("--- Setting Log Level for Connect worker at", host, "---")
			fmt.Println("making a call to", loggerListURL) // control statement print - TOREMOVE
			response, err := utilities.DoCallByHost(http.MethodPut, loggerListURL, bytes.NewBuffer(buildSetPayload()))
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				printSetResponse(response)
			}
		}
	},
}

func init() {
	LoggerSetCmd.Flags().StringVarP(&setPluginClass, "plugin-class", "", "", "plugin class to check for log level (required)")
	LoggerSetCmd.Flags().StringVarP(&setLevel, "level", "", "", "log level to be set (required)")
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

func printSetResponse(response *http.Response) {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	utilities.PrettyPrint(body)
}
