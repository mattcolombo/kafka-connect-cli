package task

import (
	"github.com/spf13/cobra"
)

var connectorName string

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "manage connector tasks",
	Long:  "Allows to manage connector tasks",
}

func init() {
	TaskCmd.PersistentFlags().StringVarP(&connectorName, "name", "n", "", "name of the connector (required)")
	TaskCmd.MarkPersistentFlagRequired("name")
	TaskCmd.AddCommand(TaskListCmd)
	TaskCmd.AddCommand(TaskGetCmd)
	TaskCmd.AddCommand(TaskRestartCmd)
}
