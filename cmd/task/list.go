package task

import (
	"fmt"
	"io/ioutil"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var showStatus, showInfo bool

var TaskListCmd = &cobra.Command{
	Use:   "list",
	Short: "short description",
	Long:  "long description",
	Run: func(cmd *cobra.Command, args []string) {
		var listURL string = buildAddress()
		fmt.Println("making a call to", listURL)
		doCall(listURL)
	},
}

func init() {
	TaskListCmd.Flags().BoolVarP(&showStatus, "show-status", "", false, "whether the command should show or not the status for each connector")
	TaskListCmd.Flags().BoolVarP(&showInfo, "show-info", "", false, "whether the command should expand or not on extra info for each connector")
}

func buildAddress() string {
	address := "http://" + utilities.ConnectConfiguration.Hostname[0] + "/connectors"
	if showStatus && showInfo {
		address += "?expand=status&expand=info"
		return address
	}
	if showStatus {
		address += "?expand=status"
	}
	if showInfo {
		address += "?expand=info"
	}
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
