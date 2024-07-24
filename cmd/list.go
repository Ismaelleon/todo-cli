package cmd

import (
	"github.com/Ismaelleon/todo-cli/tasks"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use: "list",
	Short: "Lists the tasks list",
	Run: func(cmd *cobra.Command, args []string) {
		tasks.List()
	},
}
