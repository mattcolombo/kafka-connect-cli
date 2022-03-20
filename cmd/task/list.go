package task

import (
	"fmt"
	"io/ioutil"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

//var connectorName string

var TaskListCmd = &cobra.Command{
	Use:   "list",
	Short: "task list short description",
	Long:  "task list long description",
	Run: func(cmd *cobra.Command, args []string) {
		var taskListURL string = buildListAddress()
		fmt.Println("making a call to", taskListURL) // control statement print - TOREMOVE
		doListCall(taskListURL)
	},
}

/*
func init() {
	TaskListCmd.Flags().StringVarP(&connectorName, "name", "n", "", "name of the connector to get tasks for (required)")
	TaskListCmd.MarkFlagRequired("name")
}
*/

func buildListAddress() string {
	address := "http://" + utilities.ConnectConfiguration.Hostname[0] + "/connectors/" + connectorName + "/tasks"
	return address
}

func doListCall(address string) {
	response, err := utilities.ConnectClient.Get(address)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		utilities.PrettyPrint(data)
	}
}
