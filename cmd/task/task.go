package task

import (
	"github.com/spf13/cobra"
)

var connectorName string

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "task short description",
	Long:  "task long description",
}

func init() {
	TaskCmd.PersistentFlags().StringVarP(&connectorName, "name", "n", "", "name of the connector to get tasks for (required)")
	TaskCmd.MarkPersistentFlagRequired("name")
	TaskCmd.AddCommand(TaskListCmd)
	TaskCmd.AddCommand(TaskGetCmd)
	TaskCmd.AddCommand(TaskRestartCmd)
}
