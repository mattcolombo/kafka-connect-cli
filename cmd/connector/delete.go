package connector

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

// TODO add capability to print outcome that is not JSON (see task restart); Probably this needs to become a common printer function

var ConnectorDeleteCmd = &cobra.Command{
	Use:   "delete [flags] connector_name",
	Short: "delete a connector",
	Long:  "Allows to delete a connector",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		connectorName = args[0]
		var path = fmt.Sprintf("/connectors/%s", connectorName)
		//fmt.Println("making a call to", path) // control statement print
		response, err := utilities.DoCallByPath(http.MethodDelete, path, nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			message := fmt.Sprintf("Connector %s was deleted successfully", connectorName)
			utilities.PrintEmptyBodyResponse(response, 204, message)
		}
	},
}
