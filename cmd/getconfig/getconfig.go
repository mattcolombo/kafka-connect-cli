package getconfig

import (
	"fmt"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var GetConfigCmd = &cobra.Command{
	Use:   "get-config",
	Short: "shows details of the active CLI config file",
	Long:  "shows the path of the currently active CLI config file and prints the config file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Config file (shown below) is currently loaded from path: %s\n", utilities.ConfigLoadPath)
		fmt.Println("---")
		utilities.PrettyPrintConfigYaml(utilities.ConnectConfiguration)
	},
}
