package task

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var TaskRestartCmd = &cobra.Command{
	Use:   "restart [flags] connector_name task_id",
	Short: "restarts a connector task",
	Long:  "Allows to restart a specific task for a connector",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		connectorName = args[0]
		taskID = args[1]
    validateTaskIdInput(taskID)
		var path = fmt.Sprintf("/connectors/%s/tasks/%s/restart", connectorName, taskID)
		//fmt.Println("making a call to", path) // control statement print
		response, err := utilities.DoCallByPath(http.MethodPost, path, nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			message := fmt.Sprintf("Task %s for connector %s was restarted successfully", taskID, connectorName)
			utilities.PrintEmptyBodyResponse(response, 204, message)
		}
	},
}
