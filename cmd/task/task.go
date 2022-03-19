package task

import (
	"github.com/spf13/cobra"
)

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "task short description",
	Long:  "task long description",
}

func init() {
	TaskCmd.AddCommand(TaskListCmd)
}
