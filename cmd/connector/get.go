package connector

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var configOnly, statusOnly bool

var ConnectorGetCmd = &cobra.Command{
	Use:   "get [flags] connector_name",
	Short: "shows information on a connector",
	Long:  "Shows the information on a connector, which includes configuration and status of the connector and tasks",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		connectorName = args[0]
		// check that only one of the status and cofig flags are used (if any)
		utilities.CheckMutuallyExclusive(configOnly, statusOnly, "the --status and --config flags are mutually exclusive. Please use only one.")
		var path string = buildGetPath()
		//fmt.Println("making a call to", path) // control statement print
		response, err := utilities.DoCallByPath(http.MethodGet, path, nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			utilities.PrintResponseJson(response)
		}
	},
}

func init() {
	ConnectorGetCmd.Flags().BoolVarP(&configOnly, "config-only", "c", false, "shows only the status of the connector (cannot be used together with --status-only)")
	ConnectorGetCmd.Flags().BoolVarP(&statusOnly, "status-only", "s", false, "shows only the connector configuration (cannot be used together with --config-only)")
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
