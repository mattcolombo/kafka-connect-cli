package task

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var TaskListCmd = &cobra.Command{
	Use:   "list [flags] connector_name",
	Short: "lists the tasks for a connector",
	Long:  "Allows to get a list of the tasks owned by a connector",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		connectorName = args[0]
		var path string = "/connectors/" + connectorName + "/tasks"
		//fmt.Println("making a call to", path) // control statement print
		response, err := utilities.DoCallByPath(http.MethodGet, path, nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			utilities.PrintResponseJson(response)
		}
	},
}
