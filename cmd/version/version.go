package version

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var printJson bool

var cliVersion = utilities.Version{
	Major:      1,
	Minor:      0,
	GitVersion: "v1.0.1",
	GitCommit:  "manual_build",
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "shows the CLI version and configuration used (if specified)",
	Long:  "shows the CLI version and configuration used (if specified)",
	Run: func(cmd *cobra.Command, args []string) {
		if printJson {
			byte, err := json.Marshal(&cliVersion)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			utilities.PrettyPrintJson(byte)
		} else {
			printSimpleVersion()
		}
	},
}

func init() {
	VersionCmd.Flags().BoolVarP(&printJson, "json", "j", false, "prints the version information as Json")
}

func printSimpleVersion() {
	fmt.Println("GitVersion:", cliVersion.GitVersion)
	fmt.Println("GitCommit:", cliVersion.GitCommit)
}
