package task

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var taskRestartID int

var TaskRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "task restart short description",
	Long:  "task restart long description",
	Run: func(cmd *cobra.Command, args []string) {
		var taskRestartURL string = buildRestartAddress()
		fmt.Println("making a call to", taskRestartURL) // control statement print - TOREMOVE
		doRestartCall(taskRestartURL)
	},
}

func init() {
	TaskRestartCmd.Flags().IntVarP(&taskRestartID, "id", "", 0, "ID of the task to restart (required)")
	TaskRestartCmd.MarkFlagRequired("id")
}

func buildRestartAddress() string {
	address := "http://" + utilities.ConnectConfiguration.Hostname[0] + "/connectors/" + connectorName + "/tasks/" + strconv.Itoa(taskRestartID) + "/restart"
	return address
}

func doRestartCall(address string) {
	//r := strings.NewReader("")
	response, err := utilities.ConnectClient.Post(address, "*/*", nil)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode == 204 {
			fmt.Println("Connector restarted successfully")
		}
		fmt.Println("HTTP Response:", response.StatusCode, http.StatusText(response.StatusCode))
		data, _ := ioutil.ReadAll(response.Body)
		utilities.PrettyPrint(data)
	}
}
