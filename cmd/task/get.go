package task

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var TaskGetCmd = &cobra.Command{
	Use:   "get [flags] connector_name task_id",
	Short: "shows info on a task",
	Long:  "Allows to gather information on a specific task owned by a connector",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		connectorName = args[0]
		taskID = args[1]
    validateTaskIdInput(taskID)
		var path = fmt.Sprintf("/connectors/%s/tasks/%s/status", connectorName, taskID)
		//fmt.Println("making a call to", path) // control statement print
		response, err := utilities.DoCallByPath(http.MethodGet, path, nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			utilities.PrintResponseJson(response)
		}
	},
}
