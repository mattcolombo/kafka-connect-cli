package cluster

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var showPlugins bool

var ClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Gather information on a Connect cluster",
	Long:  "Allows to gather information on a Connect cluster, and if required on the plugins installed",
}

var ClusterGet = &cobra.Command{
	Use:   "get",
	Short: "Show info about the cluster",
	Long:  "Shows the information about the cluster, and if required produces a list of the installed plugins",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--- Connect Cluster Info ---")
		getInfo("/")
		if showPlugins {
			fmt.Println("--- Installed Plugins ---")
			getInfo("/connector-plugins")
		}
	},
}

func init() {
	ClusterCmd.AddCommand(ClusterGet)
	ClusterGet.Flags().BoolVarP(&showPlugins, "show-plugins", "", false, "whether the command should show or not the list of plugins currently installed")
}

func getInfo(path string) {
	response, err := utilities.DoCallByPath(http.MethodGet, path, nil)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		utilities.PrintResponseJson(response)
	}
}
