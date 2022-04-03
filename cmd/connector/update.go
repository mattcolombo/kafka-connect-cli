package connector

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var ConnectorUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "short description",
	Long:  "long description",
	Run: func(cmd *cobra.Command, args []string) {
		var path string = "/connectors/" + connectorName + "/config"
		fmt.Println("making a call to", path) // control statement print - TOREMOVE
		response, err := utilities.DoCallByPath(http.MethodPut, path, buildUpdateRequestBody())
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

func buildUpdateRequestBody() io.Reader {
	fmt.Println("I am importing the configuration file from", connectorPath) // control statement print - TOREMOVE
	file, err := ioutil.ReadFile(connectorPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Connector config file not found!")
			os.Exit(1)
		} else {
			fmt.Println("Error while opening the connector configuration file")
			fmt.Println(err)
			os.Exit(1)
		}
	}

	/*
		Update endpoint expects to receive only the configuration, not the full JSON structure of the configuration file.
		In order to build that (since we don't know the specific structure of the connector being updated) we first unmarshal it into
		a structure and extract the "config" structure. This is extracted as a map, and then this map is Marshaled again as JSON to be
		able to pass it to the PUT Rest endpoint

		NOTE: the structure is shared by different commands so is defined in connector.go in the common variables section
	*/
	// unmarshaling the config file and extracting the "config" part into a map --> Config map[string]string `json:"config"`
	var jsonConfig connectConfig
	json.Unmarshal(file, &jsonConfig)
	if err != nil {
		fmt.Printf("Error while extracting the connector configuration %s\n", err)
		os.Exit(1)
	}

	// marshaling the map obtained in the previous step to build the request body
	jsonData, err := json.Marshal(jsonConfig.Config)
	if err != nil {
		fmt.Printf("Error while building the request body %s\n", err)
		os.Exit(1)
	}
	utilities.PrettyPrintJson(jsonData) //TOREMOVE control statement
	return bytes.NewBuffer(jsonData)
}
