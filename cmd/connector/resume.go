package connector

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

// TODO add capability to print outcome that is not JSON (see task restart); Probably this needs to become a common printer function

var ConnectorResumeCmd = &cobra.Command{
	Use:   "resume [flags] connector_name",
	Short: "resume a connector",
	Long:  "Allows to resume processing for a connector that was previously paused",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		connectorName = args[0]
		var path string = "/connectors/" + connectorName + "/resume"
		//fmt.Println("making a call to", path) // control statement print
		response, err := utilities.DoCallByPath(http.MethodPut, path, nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			message := fmt.Sprintf("Connector %s was resumed successfully", connectorName)
			utilities.PrintEmptyBodyResponse(response, 202, message)
		}
	},
}
