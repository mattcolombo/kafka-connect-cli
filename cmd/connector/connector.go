package connector

import (
	"github.com/spf13/cobra"
)

// common variables that will be used in multiple commands
var connectorName, connectorConfigPath string

var ConnectorCmd = &cobra.Command{
	Use:   "connector",
	Short: "short description",
	Long:  "long description",
}

func init() {
	ConnectorCmd.AddCommand(ConnectorListCmd)
	ConnectorCmd.AddCommand(ConnectorGetCmd)
}
