package logger

import (
	"fmt"
	"io/ioutil"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var showStatus, showInfo bool

var LoggerListCmd = &cobra.Command{
	Use:   "list",
	Short: "logger list short description",
	Long:  "logger list long description",
	Run: func(cmd *cobra.Command, args []string) {
		for _, host := range utilities.ConnectConfiguration.Hostname {
			var loggerListURL string = buildAddress(host)
			fmt.Println("--- Loggers Info for", host, "---")
			fmt.Println("making a call to", loggerListURL)
			doCall(loggerListURL)
		}
	},
}

func buildAddress(host string) string {
	address := "http://" + host + "/admin/loggers"
	return address
}

func doCall(address string) {
	response, err := utilities.ConnectClient.Get(address)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		utilities.PrettyPrint(data)
	}
}
