package cmd

import (
	"github.com/Ismaelleon/todo-cli/tasks"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks.Add(args)
	},
}
