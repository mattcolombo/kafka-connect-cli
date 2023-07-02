package connector

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var includeTasks, onlyFailed bool

var ConnectorRestartCmd = &cobra.Command{
	Use:   "restart [flags] connector_name",
	Short: "restart a connector",
	Long:  "Allows to restart a connector",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		connectorName = args[0]
		var path string = buildRestartPath()
		//fmt.Println("making a call to", path) // control statement print
		response, err := utilities.DoCallByPath(http.MethodPost, path, nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			message := fmt.Sprintf("Connector %s was restarted successfully", connectorName)
			utilities.PrintEmptyBodyResponse(response, 204, message)
		}
	},
}

func init() {
	ConnectorRestartCmd.Flags().BoolVarP(&includeTasks, "include-tasks", "", false, "restart also the connector tasks")
	ConnectorRestartCmd.Flags().BoolVarP(&onlyFailed, "failed-only", "", false, "restart only the failed instances")
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
