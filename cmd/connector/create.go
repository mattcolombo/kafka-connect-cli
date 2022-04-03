package connector

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var connectorPath string

var ConnectorCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "short description",
	Long:  "long description",
	Run: func(cmd *cobra.Command, args []string) {
		var path string = "/connectors"
		fmt.Println("making a call to", path) // control statement print - TOREMOVE
		response, err := utilities.DoCallByPath(http.MethodPost, path, buildCreateRequestBody())
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			utilities.PrintResponseJson(response)
		}
	},
}

func init() {
	ConnectorCreateCmd.Flags().StringVarP(&connectorPath, "config-path", "", "", "path to the connector configuration file (required)")
	ConnectorCreateCmd.MarkFlagRequired("config-path")
}

func buildCreateRequestBody() io.Reader {
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

	return bytes.NewBuffer(file)
}
