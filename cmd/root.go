package cmd

import (
	"fmt"
	"os"

	"github.com/mattcolombo/kafka-connect-cli/cmd/cluster"
	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var ConnectConfiguration utilities.Configuration
var configPath string

var rootCmd = &cobra.Command{
	Use:   "kafka-connect-cli",
	Short: "short description",
	Long:  "long description",
}

func init() {
	// adding the flag required for getting the path to the configuration file. This is required always
	rootCmd.PersistentFlags().StringVar(&configPath, "connect-config", "", "the path to the configuration file containing the information about the Connect target (required)")
	rootCmd.MarkPersistentFlagRequired("connect-config")

	// adding the subcommands required by the CLI tool
	rootCmd.AddCommand(cluster.Cluster)
	//rootCmd.AddCommand(goodbye.Bye)
	//rootCmd.AddCommand(time.Time)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// importing the configuration for the whole CLI to use during execution
	utilities.ConnectConfiguration = utilities.ImportConfig(configPath)
	fmt.Println(utilities.ConnectConfiguration.Hostname[0])
}
