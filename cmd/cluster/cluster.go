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
	Short: "short description",
	Long:  "long description",
}

var ClusterGet = &cobra.Command{
	Use:   "get",
	Short: "short description",
	Long:  "long description",
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
