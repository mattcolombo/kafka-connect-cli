package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
			var loggerListURL string = buildSetAddress(host)
			fmt.Println("--- Loggers Info for Connect worker at", host, "---")
			fmt.Println("making a call to", loggerListURL) // control statement print - TOREMOVE
			doSetCall(loggerListURL)
		}
	},
}

func init() {
	LoggerSetCmd.Flags().StringVarP(&setPluginClass, "plugin-class", "", "", "plugin class to check for log level (required)")
	LoggerSetCmd.Flags().StringVarP(&setLevel, "level", "", "", "log level to be set (required)")
	LoggerSetCmd.MarkFlagRequired("plugin-class")
	LoggerSetCmd.MarkFlagRequired("level")
}

func buildSetAddress(host string) string {
	address := "http://" + host + "/admin/loggers/" + setPluginClass
	return address
}

func doSetCall(address string) {
	payload, err := json.Marshal(map[string]interface{}{"level": setLevel})
	if err != nil {
		fmt.Printf("JSON build failed with error %s\n", err)
	}

	request, err := http.NewRequest(http.MethodPut, address, bytes.NewBuffer(payload))
	request.Header.Set("Content-Type", "application/json")
	if err != nil {
		fmt.Printf("Creation of request failed with error %s\n", err)
	}

	response, err := utilities.ConnectClient.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		utilities.PrettyPrint(data)
	}
}
