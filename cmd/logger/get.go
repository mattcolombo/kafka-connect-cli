package logger

import (
	"fmt"
	"io/ioutil"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var getPluginClass string

var LoggerGetCmd = &cobra.Command{
	Use:   "get",
	Short: "logger get short description",
	Long:  "logger get long description",
	Run: func(cmd *cobra.Command, args []string) {
		for _, host := range utilities.ConnectConfiguration.Hostname {
			var loggerListURL string = buildGetAddress(host)
			fmt.Println("--- Loggers Info for Connect worker at", host, "---")
			fmt.Println("making a call to", loggerListURL) // control statement print - TOREMOVE
			doGetCall(loggerListURL)
		}
	},
}

func init() {
	LoggerGetCmd.Flags().StringVarP(&getPluginClass, "plugin-class", "", "", "plugin class to check for log level (required)")
	LoggerGetCmd.MarkFlagRequired("plugin-class")
}

func buildGetAddress(host string) string {
	address := "http://" + host + "/admin/loggers/" + getPluginClass
	return address
}

func doGetCall(address string) {
	response, err := utilities.ConnectClient.Get(address)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		utilities.PrettyPrint(data)
	}
}
