package connector

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

// common variables that will be used in multiple commands
var connectorName, connectorPath string

// structure used to extract data from the configuration file for the connector (the config part for update and the connector class for validation)
type connectConfig struct {
	Name   string            `json:"name"`
	Config map[string]string `json:"config"`
}

var ConnectorCmd = &cobra.Command{
	Use:   "connector",
	Short: "short description",
	Long:  "long description",
}

func init() {
	ConnectorCmd.AddCommand(ConnectorListCmd)
	ConnectorCmd.AddCommand(ConnectorGetCmd)
	ConnectorCmd.AddCommand(ConnectorCreateCmd)
	ConnectorCmd.AddCommand(ConnectorUpdateCmd)
	ConnectorCmd.AddCommand(ConnectorDeleteCmd)
	ConnectorCmd.AddCommand(ConnectorPauseCmd)
	ConnectorCmd.AddCommand(ConnectorResumeCmd)
	ConnectorCmd.AddCommand(ConnectorRestartCmd)
}

/*
Common functions around extracting and building the request bodies

These functions are not added to the utilities since they are only used within
this package so do not need to be available to the rest of the CLI.
They may be moved in the future if required while refactoring.
*/
func extractRequestBody(filePath string) []byte {
	fmt.Println("I am importing the configuration file from", filePath) // control statement print - TOREMOVE
	file, err := ioutil.ReadFile(filePath)
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
	return file
}

func extractConnectorConfig(file []byte) []byte {
	/*
		Update endpoint expects to receive only the configuration, not the full JSON structure of the configuration file.
		In order to build that (since we don't know the specific structure of the connector being updated) we first unmarshal it into
		a structure and extract the "config" structure. This is extracted as a map, and then this map is Marshaled again as JSON to be
		able to pass it to the PUT Rest endpoint

		NOTE: the structure is shared by different commands so is defined in connector.go in the common variables section
	*/
	// unmarshaling the config file and extracting the "config" part into a map --> Config map[string]string `json:"config"`
	var jsonConfig connectConfig
	err := json.Unmarshal(file, &jsonConfig)
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
	return jsonData
}
