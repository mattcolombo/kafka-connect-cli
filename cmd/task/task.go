package task

import (
	"github.com/spf13/cobra"
)

var connectorName string
var taskID string

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "manage connector tasks",
	Long:  "Allows to manage connector tasks",
}

func init() {
	TaskCmd.AddCommand(TaskListCmd)
	TaskCmd.AddCommand(TaskGetCmd)
	TaskCmd.AddCommand(TaskRestartCmd)
}
