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
	Short: "Gather information on the Connect cluster",
	Long:  "Gather information on the Connect cluster specified in the configuration file used; if required also lists the plugins installed",
}

var ClusterGet = &cobra.Command{
	Use:   "get",
	Short: "show info about the cluster",
	Long:  "Shows the information about the cluster; can also produce a list of the installed plugins",
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
	ClusterGet.Flags().BoolVarP(&showPlugins, "show-plugins", "", false, "produces the list of plugins currently installed along with the cluster info")
}

func getInfo(path string) {
	response, err := utilities.DoCallByPath(http.MethodGet, path, nil)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		utilities.PrintResponseJson(response)
	}
}
