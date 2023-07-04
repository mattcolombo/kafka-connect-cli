package cmd

import (
	"fmt"
	"os"

	"github.com/mattcolombo/kafka-connect-cli/cmd/cluster"
	"github.com/mattcolombo/kafka-connect-cli/cmd/connector"
	"github.com/mattcolombo/kafka-connect-cli/cmd/getconfig"
	"github.com/mattcolombo/kafka-connect-cli/cmd/logger"
	"github.com/mattcolombo/kafka-connect-cli/cmd/task"
	"github.com/mattcolombo/kafka-connect-cli/cmd/version"
	"github.com/spf13/cobra"
)

//var version = "1.1.0"

var rootCmd = &cobra.Command{
	Use: "kconnect-cli",
	//Version: version,
	Short: "command line tool to manage a Kafka Connect installation",
	Long: `A comprehensive command line tool to manage a Kafka Connect installation. 
Allows to gather information about the cluster, connectors, tasks loggers and manage them.
---
Requires a configuration file either selected through an environment variable "CONNECTCFG" or located in the current working directory.
---
Further details and documentation can be found at https://github.com/mattcolombo/kafka-connect-cli`,
}

func init() {
	// adding the subcommands required by the CLI tool
	rootCmd.AddCommand(cluster.ClusterCmd)
	rootCmd.AddCommand(connector.ConnectorCmd)
	rootCmd.AddCommand(task.TaskCmd)
	rootCmd.AddCommand(logger.LoggerCmd)
	rootCmd.AddCommand(version.VersionCmd)
	rootCmd.AddCommand(getconfig.GetConfigCmd)

	// disabled the default completion command as it is not going to be necessary
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
