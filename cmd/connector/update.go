package connector

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var ConnectorUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "short description",
	Long:  "long description",
	Run: func(cmd *cobra.Command, args []string) {
		connectorConfiguration := extractRequestBody(connectorPath)
		configData := extractConnectorConfig(connectorConfiguration)
		requestBody := bytes.NewBuffer(configData)
		var path string = "/connectors/" + connectorName + "/config"
		fmt.Println("making a call to", path) // control statement print - TOREMOVE
		response, err := utilities.DoCallByPath(http.MethodPut, path, requestBody)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			utilities.PrintResponseJson(response)
		}
	},
}

func init() {
	ConnectorUpdateCmd.Flags().StringVarP(&connectorName, "name", "n", "", "name of the connector to update (required)")
	ConnectorUpdateCmd.MarkFlagRequired("name")
	ConnectorUpdateCmd.Flags().StringVarP(&connectorPath, "config-path", "", "", "path to the connector configuration file (required)")
	ConnectorUpdateCmd.MarkFlagRequired("config-path")
}
