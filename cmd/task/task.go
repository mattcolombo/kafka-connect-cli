package task

import (
	"github.com/spf13/cobra"
)

var connectorName string

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage connector tasks",
	Long:  "Allows to manage connector tasks",
}

func init() {
	TaskCmd.PersistentFlags().StringVarP(&connectorName, "name", "n", "", "name of the connector to get tasks for (required)")
	TaskCmd.MarkPersistentFlagRequired("name")
	TaskCmd.AddCommand(TaskListCmd)
	TaskCmd.AddCommand(TaskGetCmd)
	TaskCmd.AddCommand(TaskRestartCmd)
}
