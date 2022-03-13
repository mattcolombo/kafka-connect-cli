package cmd

import (
	"fmt"
	"os"

	"github.com/mattcolombo/kafka-connect-cli/cmd/cluster"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kafka-connect-cli",
	Short: "short description",
	Long:  "long description",
}

func init() {
	rootCmd.AddCommand(cluster.Cluster)
	//rootCmd.AddCommand(goodbye.Bye)
	//rootCmd.AddCommand(time.Time)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
