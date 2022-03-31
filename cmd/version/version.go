package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version string = "kafka-connect-cli : {Major: 1, Minor: 0, Full version: 1.0.0}"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "version short description",
	Long:  "version long description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
