package cmd

import (
	"github.com/Ismaelleon/todo-cli/tasks"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use: "set",
	Short: "Set the task state (done or not done)",
	Run: func(cmd *cobra.Command, args []string) {
		tasks.Set(args)
	},
}
