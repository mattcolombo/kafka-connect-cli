package connector

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var showStatus, showInfo bool

var ConnectorListCmd = &cobra.Command{
	Use:   "list",
	Short: "list the connectors on the cluster",
	Long:  "Produces a list of all the connectors currently running on the cluster",
	Run: func(cmd *cobra.Command, args []string) {
		var path string = buildListPath()
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
	ConnectorListCmd.Flags().BoolVarP(&showStatus, "show-status", "", false, "show also the status for each connector")
	ConnectorListCmd.Flags().BoolVarP(&showInfo, "show-info", "", false, "expand extra info for each connector")
}

func buildListPath() string {
	path := "/connectors"
	if showStatus && showInfo {
		path += "?expand=status&expand=info"
		return path
	}
	if showStatus {
		path += "?expand=status"
	}
	if showInfo {
		path += "?expand=info"
	}
	return path
}
