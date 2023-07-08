package task

import (
	"fmt"
	"os"
	"strconv"

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

func validateTaskIdInput(id string) {
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("task_id must be a digit")
		os.Exit(1)
	}
}
