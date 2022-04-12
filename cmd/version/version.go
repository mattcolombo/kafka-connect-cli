package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string = "kconnect-cli : {Major: 0, Minor: 0, Full version: 0.0.1}"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "shows the CLI version",
	Long:  "Shows the CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
