package task

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
		var path string = "/connectors/" + connectorName + "/tasks/" + strconv.Itoa(taskGetID) + "/status"
		fmt.Println("making a call to", path) // control statement print - TOREMOVE
		response, err := utilities.DoCallByPath(http.MethodGet, path, nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			printGetResponse(response)
		}
	},
}

func init() {
	TaskGetCmd.Flags().IntVarP(&taskGetID, "id", "", 0, "ID of the task to check (required)")
	TaskGetCmd.MarkFlagRequired("id")
}

func printGetResponse(response *http.Response) {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	utilities.PrettyPrint(body)
}
