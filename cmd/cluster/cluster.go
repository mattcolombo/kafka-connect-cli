package cluster

import (
	"fmt"
	"io/ioutil"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var showPlugins bool

//var client *http.Client = utilities.CreateClient(utilities.ConnectConfiguration)

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
		getConnectInfo()
		if showPlugins {
			fmt.Println("--- Installed Plugins ---")
			getConnectPlugins()
		}
	},
}

func init() {
	ClusterCmd.AddCommand(ClusterGet)
	ClusterGet.Flags().BoolVarP(&showPlugins, "show-plugins", "", false, "whether the command should show or not the list of plugins currently installed")
}

func getConnectInfo() {
	address := "http://" + utilities.ConnectConfiguration.Hostname[0] + "/"
	response, err := utilities.ConnectClient.Get(address)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		utilities.PrettyPrint(data)
	}
}

func getConnectPlugins() {
	address := "http://" + utilities.ConnectConfiguration.Hostname[0] + "/connector-plugins"
	response, err := utilities.ConnectClient.Get(address)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		utilities.PrettyPrint(data)
	}
}
