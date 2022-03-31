package connector

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var showStatus, showInfo bool

var ConnectorListCmd = &cobra.Command{
	Use:   "list",
	Short: "short description",
	Long:  "long description",
	Run: func(cmd *cobra.Command, args []string) {
		var path string = buildPath()
		fmt.Println("making a call to", path) // control statement print - TOREMOVE
		response, err := utilities.DoCallByPath(http.MethodGet, path, nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			printListResponse(response)
		}
	},
}

func init() {
	ConnectorListCmd.Flags().BoolVarP(&showStatus, "show-status", "", false, "whether the command should show or not the status for each connector")
	ConnectorListCmd.Flags().BoolVarP(&showInfo, "show-info", "", false, "whether the command should expand or not on extra info for each connector")
}

func printListResponse(response *http.Response) {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	utilities.PrettyPrint(body)
}

func buildPath() string {
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
