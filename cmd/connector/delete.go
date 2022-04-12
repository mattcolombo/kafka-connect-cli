package connector

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

// TODO add capability to print outcome that is not JSON (see task restart); Probably this needs to become a common printer function

var ConnectorDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a connector",
	Long:  "Allows to delete a connector",
	Run: func(cmd *cobra.Command, args []string) {
		var path string = "/connectors/" + connectorName
		fmt.Println("making a call to", path) // control statement print - TOREMOVE
		response, err := utilities.DoCallByPath(http.MethodDelete, path, nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			message := fmt.Sprintf("Connector %s was deleted successfully", connectorName)
			utilities.PrintEmptyBodyResponse(response, 204, message)
		}
	},
}

func init() {
	ConnectorDeleteCmd.Flags().StringVarP(&connectorName, "name", "n", "", "name of the connector to delete (required)")
	ConnectorDeleteCmd.MarkFlagRequired("name")
}
