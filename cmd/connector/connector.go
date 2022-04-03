package connector

import (
	"github.com/spf13/cobra"
)

// common variables that will be used in multiple commands
var connectorName, connectorPath string

// structure used to extract data from the configuration file for the connector (the config part for update and the connector class for validation)
type connectConfig struct {
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
}
