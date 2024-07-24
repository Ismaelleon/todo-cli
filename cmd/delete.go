package cmd

import (
	"github.com/Ismaelleon/todo-cli/tasks"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use: "delete",
	Short: "Deletes a task by its index",
	Run: func(cmd *cobra.Command, args []string) {
		tasks.Delete(args)
	},
}
