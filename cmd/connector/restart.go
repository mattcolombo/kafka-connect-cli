package connector

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var includeTasks, onlyFailed bool

var ConnectorRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "short description",
	Long:  "long description",
	Run: func(cmd *cobra.Command, args []string) {
		var path string = buildRestartPath()
		fmt.Println("making a call to", path) // control statement print - TOREMOVE
		response, err := utilities.DoCallByPath(http.MethodPost, path, nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			message := fmt.Sprintf("Connector %s was restarted successfully", connectorName)
			utilities.PrintEmptyBodyResponse(response, 204, message)
			//fmt.Println(response.Status)
			//utilities.PrintResponseJson(response)
		}
	},
}

func init() {
	ConnectorRestartCmd.Flags().StringVarP(&connectorName, "name", "n", "", "name of the connector to show (required)")
	ConnectorRestartCmd.MarkFlagRequired("name")
	ConnectorRestartCmd.Flags().BoolVarP(&includeTasks, "include-tasks", "", false, "whether to restart also the connector tasks")
	ConnectorRestartCmd.Flags().BoolVarP(&onlyFailed, "failed-only", "", false, "whether to restart only the failed instances")
}

func buildRestartPath() string {
	path := "/connectors/" + connectorName + "/restart"
	if includeTasks && onlyFailed {
		path += "?includeTasks=true&onlyFailed=true"
		return path
	}
	if includeTasks {
		path += "?includeTasks=true"
	}
	if onlyFailed {
		path += "?onlyFailed=true"
	}
	return path
}
