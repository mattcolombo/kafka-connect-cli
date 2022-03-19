package cmd

import (
	"fmt"
	"os"

	"github.com/mattcolombo/kafka-connect-cli/cmd/cluster"
	"github.com/mattcolombo/kafka-connect-cli/cmd/connector"
	"github.com/spf13/cobra"
)

//var configPath string

var rootCmd = &cobra.Command{
	Use:   "kafka-connect-cli",
	Short: "short description",
	Long:  "long description - remember to add that the file needs to be added as an Environment variable",
}

func init() {
	// adding the flag required for getting the path to the configuration file. This is required always
	//rootCmd.PersistentFlags().StringVar(&configPath, "connect-config", "", "the path to the configuration file containing the information about the Connect target (required)")
	//rootCmd.MarkPersistentFlagRequired("connect-config")

	// adding the subcommands required by the CLI tool
	rootCmd.AddCommand(cluster.ClusterCmd)
	rootCmd.AddCommand(connector.ConnectorCmd)
}

func Execute() {
	//os.Setenv("CONNECTCFG", configPath)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
