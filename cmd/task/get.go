package task

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var taskGetID int

var TaskGetCmd = &cobra.Command{
	Use:   "get",
	Short: "shows info on a task",
	Long:  "Allows to gather information on a specific task owned by a connector",
	Run: func(cmd *cobra.Command, args []string) {
		var path string = "/connectors/" + connectorName + "/tasks/" + strconv.Itoa(taskGetID) + "/status"
		fmt.Println("making a call to", path) // control statement print - TOREMOVE
		response, err := utilities.DoCallByPath(http.MethodGet, path, nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			utilities.PrintResponseJson(response)
		}
	},
}

func init() {
	TaskGetCmd.Flags().IntVarP(&taskGetID, "id", "", 0, "ID of the task to check (required)")
	TaskGetCmd.MarkFlagRequired("id")
}
