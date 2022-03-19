package cmd

import (
	"fmt"
	"os"

	"github.com/mattcolombo/kafka-connect-cli/cmd/cluster"
	"github.com/mattcolombo/kafka-connect-cli/cmd/connector"
	"github.com/mattcolombo/kafka-connect-cli/cmd/logger"
	"github.com/mattcolombo/kafka-connect-cli/cmd/task"
	"github.com/spf13/cobra"
)

//var configPath string

var rootCmd = &cobra.Command{
	Use:   "kafka-connect-cli",
	Short: "short description",
	Long:  "long description - remember to add that the file needs to be added as an Environment variable",
}

func init() {
	// adding the subcommands required by the CLI tool
	rootCmd.AddCommand(cluster.ClusterCmd)
	rootCmd.AddCommand(connector.ConnectorCmd)
	rootCmd.AddCommand(task.TaskCmd)
	rootCmd.AddCommand(logger.LoggerCmd)

	// disabled the default completion command as it is not going to be necessary
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
