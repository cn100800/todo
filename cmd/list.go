package cmd

import (
	"github.com/freecracy/todo/task"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `A longer description application.`,
	Run: func(cmd *cobra.Command, args []string) {
		task.ListTask()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
