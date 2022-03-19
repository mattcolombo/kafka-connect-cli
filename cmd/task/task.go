package task

import (
	"github.com/spf13/cobra"
)

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "short description",
	Long:  "long description",
}

func init() {
	TaskCmd.AddCommand(TaskListCmd)
}
