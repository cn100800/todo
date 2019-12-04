package cmd

import (
	"github.com/freecracy/todo/task"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long:  `添加任务`,
	Run: func(cmd *cobra.Command, args []string) {
		task.AddTask(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
