package cmd

import (
	"github.com/freecracy/todo/task"
	"github.com/spf13/cobra"
)

// reopenCmd represents the reopen command
var reopenCmd = &cobra.Command{
	Use:   "reopen",
	Short: "task reopen job",
	Long:  `重新打开任务`,
	Run: func(cmd *cobra.Command, args []string) {
		task.ReopenTask(args[0])
	},
}

func init() {
	rootCmd.AddCommand(reopenCmd)

	// Here you will define your flags and configuration settings.

	// and all subcommands, e.g.:
	// reopenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// is called directly, e.g.:
	// reopenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
