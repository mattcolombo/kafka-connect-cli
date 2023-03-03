package task

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var taskRestartID int

var TaskRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "restarts a connector task",
	Long:  "Allows to restart a specific task for a connector",
	Run: func(cmd *cobra.Command, args []string) {
		var path string = "/connectors/" + connectorName + "/tasks/" + strconv.Itoa(taskRestartID) + "/restart"
		//fmt.Println("making a call to", path) // control statement print
		response, err := utilities.DoCallByPath(http.MethodPost, path, nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			message := fmt.Sprintf("Task %d for connector %s was restarted successfully", taskRestartID, connectorName)
			utilities.PrintEmptyBodyResponse(response, 204, message)
		}
	},
}

func init() {
	TaskRestartCmd.Flags().IntVarP(&taskRestartID, "id", "", 0, "ID of the task to restart (required)")
	TaskRestartCmd.MarkFlagRequired("id")
}
