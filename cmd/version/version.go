package version

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var printJson bool
var MajorVersion = "--manual_build--"
var MinorVersion = "--manual_build--"
var GitVersion = "--manual_build--"
var GitHash = "--manual_build--"
var BuildDate = "--manual_build--"
var GoVersion = "--manual_build--"

var cliVersion = utilities.Version{
	Major:      MajorVersion,
	Minor:      MinorVersion,
	GitVersion: GitVersion,
	GitCommit:  GitHash,
	BuildDate:  BuildDate,
	GoVersion:  GoVersion,
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "shows the CLI version",
	Long:  "shows the short CLI version; allows JSON extended print if specified",
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
	fmt.Println("Git Version:", cliVersion.GitVersion)
	fmt.Println("Git Commit:", cliVersion.GitCommit)
	fmt.Println("Build Date:", cliVersion.BuildDate)
}
