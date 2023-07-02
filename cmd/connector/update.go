package connector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var ConnectorUpdateCmd = &cobra.Command{
	Use:   "update [flags] connector_name",
	Short: "update a connector configuration",
	Long:  "Allows to update a connector configuration from an updated configuration file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		connectorName = args[0]
		connectorConfiguration := extractRequestBody(connectorPath)
		connectorNameFromConfig := extractConnectorName(connectorConfiguration)
		if connectorName != connectorNameFromConfig {
			fmt.Println("The connector specified does not match the name in the configuration file. Please make sure you are updating the right connector")
			fmt.Println("Requested:", connectorName, "/ In config file:", connectorNameFromConfig)
			os.Exit(1)
		}
		configData := extractConnectorConfig(connectorConfiguration)
		requestBody := bytes.NewBuffer(configData)
		var path string = "/connectors/" + connectorName + "/config"
		//fmt.Println("making a call to", path) // control statement print
		response, err := utilities.DoCallByPath(http.MethodPut, path, requestBody)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			utilities.PrintResponseJson(response)
		}
	},
}

func init() {
	ConnectorUpdateCmd.Flags().StringVarP(&connectorPath, "config-file", "f", "", "path to the connector JSON configuration file (required)")
	ConnectorUpdateCmd.MarkFlagRequired("config-file")
}

func extractConnectorName(file []byte) string {
	var jsonConfig connectConfig
	json.Unmarshal(file, &jsonConfig)
	return jsonConfig.Name
}
