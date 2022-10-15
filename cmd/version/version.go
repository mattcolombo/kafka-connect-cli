package version

import (
	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var version = []byte(`{
	"Major": "0", 
	"Minor": "2", 
	"GitVersion": "---", 
	"GitCommit": "---"}`)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "shows the CLI version",
	Long:  "Shows the CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		utilities.PrettyPrintJson(version)
	},
}
