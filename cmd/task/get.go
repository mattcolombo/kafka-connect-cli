package task

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var taskGetID int

var TaskGetCmd = &cobra.Command{
	Use:   "get",
	Short: "task list short description",
	Long:  "task list long description",
	Run: func(cmd *cobra.Command, args []string) {
		var taskGetURL string = buildGetAddress()
		fmt.Println("making a call to", taskGetURL) // control statement print - TOREMOVE
		doGetCall(taskGetURL)
	},
}

func init() {
	TaskGetCmd.Flags().IntVarP(&taskGetID, "id", "", 0, "ID of the task to check (required)")
	TaskGetCmd.MarkFlagRequired("id")
}

func buildGetAddress() string {
	address := "http://" + utilities.ConnectConfiguration.Hostname[0] + "/connectors/" + connectorName + "/tasks/" + strconv.Itoa(taskGetID) + "/status"
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
