package cmd

import (
	"github.com/freecracy/todo/task"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "task list",
	Long:  `当前未完成任务`,
	Run: func(cmd *cobra.Command, args []string) {
		task.ListTask()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
