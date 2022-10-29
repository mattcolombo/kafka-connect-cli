package version

import (
	"fmt"
	"strings"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var showCliConfig bool

var version = []byte(`{
	"Major": "1", 
	"Minor": "0", 
	"GitVersion": "v1.0.0", 
	"GitCommit": "manual_bild"}`)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "shows the CLI version and configuration used (if specified)",
	Long:  "shows the CLI version and configuration used (if specified)",
	Run: func(cmd *cobra.Command, args []string) {
		utilities.PrettyPrintJson(version)
		if showCliConfig {
			printCliCfg()
		}
	},
}

func init() {
	VersionCmd.Flags().BoolVarP(&showCliConfig, "show-cli-config", "", false, "prints the location and main URL for the configuration file being loaded")
}

func printCliCfg() {
	fmt.Println("--- Current CLI Configuration ---")
	fmt.Printf("Configuration for the CLI is being loaded from path: %s\n", utilities.ConfigLoadPath)
	fmt.Printf("The main URL currently in use is <%s> with protocol %s\n", utilities.ConnectConfiguration.Hostnames[0], strings.ToUpper(utilities.ConnectConfiguration.Protocol))
}
