package connector

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var configOnly, statusOnly bool

var ConnectorGetCmd = &cobra.Command{
	Use:   "get",
	Short: "short description",
	Long:  "long description",
	Run: func(cmd *cobra.Command, args []string) {
		// check that only one of the status and cofig flags are used (if any)
		utilities.CheckMutuallyExclusive(configOnly, statusOnly, "the --status and --config flags are mutually exclusive. Please use only one.")
		var path string = buildGetPath()
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
	ConnectorGetCmd.Flags().StringVarP(&connectorName, "name", "n", "", "name of the connector to show (required)")
	ConnectorGetCmd.MarkFlagRequired("name")
	ConnectorGetCmd.Flags().BoolVarP(&configOnly, "config-only", "c", false, "shows the status of the connector (cannot be used with --status-only)")
	ConnectorGetCmd.Flags().BoolVarP(&statusOnly, "status-only", "s", false, "shows the connector configuration (cannot be used with --config-only)")
}

func buildGetPath() string {
	var path string = "/connectors/" + connectorName
	if statusOnly {
		path += "/status"
	} else if configOnly {
		path += "/config"
	}
	return path
}
