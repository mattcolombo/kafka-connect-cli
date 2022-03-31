package logger

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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
			var loggerListURL string = host + "/admin/loggers/" + getPluginClass
			fmt.Println("--- Getting Log Level for Connect worker at", host, "---")
			fmt.Println("making a call to", loggerListURL) // control statement print - TOREMOVE
			response, err := utilities.DoCallByHost(http.MethodGet, loggerListURL, nil)
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				printGetResponse(response)
			}
		}
	},
}

func init() {
	LoggerGetCmd.Flags().StringVarP(&getPluginClass, "plugin-class", "", "", "plugin class to check for log level (required)")
	LoggerGetCmd.MarkFlagRequired("plugin-class")
}

func printGetResponse(response *http.Response) {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	utilities.PrettyPrint(body)
}
