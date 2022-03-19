package connector

import (
	"github.com/spf13/cobra"
)

var ConnectorCmd = &cobra.Command{
	Use:   "connector",
	Short: "short description",
	Long:  "long description",
}

func init() {
	ConnectorCmd.AddCommand(ListCmd)
}
