package cluster

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var showPlugins bool
var client *http.Client = utilities.CreateClient(utilities.ConnectConfiguration)

var Cluster = &cobra.Command{
	Use:   "cluster",
	Short: "short description",
	Long:  "long description",
	Run: func(cmd *cobra.Command, args []string) {
		getConnectInfo()
		if showPlugins {
			getConnectPlugins()
		}
	},
}

func init() {
	Cluster.Flags().BoolVarP(&showPlugins, "show-plugins", "", false, "whether the command should show or not the list of plugins currently installed")
}

func getConnectInfo() {
	address := "http://" + utilities.ConnectConfiguration.Hostname + "/"
	response, err := client.Get(address)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		utilities.PrettyPrint(data)
	}
}

func getConnectPlugins() {
	address := "http://" + utilities.ConnectConfiguration.Hostname + "/connector-plugins"
	response, err := client.Get(address)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		utilities.PrettyPrint(data)
	}
}
