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
	Short: "task restart short description",
	Long:  "task restart long description",
	Run: func(cmd *cobra.Command, args []string) {
		var path string = "/connectors/" + connectorName + "/tasks/" + strconv.Itoa(taskRestartID) + "/restart"
		fmt.Println("making a call to", path) // control statement print - TOREMOVE
		response, err := utilities.DoCallByPath(http.MethodPost, path, nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			printRestartResponse(response)
		}
	},
}

func init() {
	TaskRestartCmd.Flags().IntVarP(&taskRestartID, "id", "", 0, "ID of the task to restart (required)")
	TaskRestartCmd.MarkFlagRequired("id")
}

// TODO refactor this in a common printer function since many connectors will require it to be implemented

func printRestartResponse(response *http.Response) {
	defer response.Body.Close()

	if response.StatusCode == 204 {
		fmt.Println("Connector restarted successfully")
	}
	fmt.Println("Connect responds:", response.StatusCode, http.StatusText(response.StatusCode))
}
