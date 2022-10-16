package connector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var validate bool

var ConnectorCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a connector",
	Long:  "Allows to create a connector from a configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		var response *http.Response
		var err error
		connectorConfiguration := extractRequestBody(connectorPath)
		if validate {
			response, err = doValidateCall(connectorConfiguration)
		} else {
			response, err = doCreateCall(connectorConfiguration)
		}
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			utilities.PrintResponseJson(response)
		}
	},
}

func init() {
	ConnectorCreateCmd.Flags().StringVarP(&connectorPath, "config-path", "", "", "path to the connector JSON configuration file (required)")
	ConnectorCreateCmd.MarkFlagRequired("config-path")
	ConnectorCreateCmd.Flags().BoolVarP(&validate, "validate", "", false, "validates the connector configuration; connector is NOT created")
}

func doCreateCall(configFile []byte) (*http.Response, error) {
	var path string = "/connectors"
	//fmt.Println("making a call to", path) // control statement print
	requestBody := bytes.NewBuffer(configFile)
	return utilities.DoCallByPath(http.MethodPost, path, requestBody)
}

func doValidateCall(configFile []byte) (*http.Response, error) {
	pluginName := extractPluginType(configFile)
	var path string = "/connector-plugins/" + pluginName + "/config/validate"
	//fmt.Println("making a call to", path) // control statement print
	configData := extractConnectorConfig(configFile)
	requestBody := bytes.NewBuffer(configData)
	return utilities.DoCallByPath(http.MethodPut, path, requestBody)
}

func extractPluginType(file []byte) string {
	var jsonConfig connectConfig
	json.Unmarshal(file, &jsonConfig)
	pluginClass := jsonConfig.Config["connector.class"]
	pluginName := strings.Split(pluginClass, ".")
	return pluginName[len(pluginName)-1]
}
