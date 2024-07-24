package cmd

import (
	"fmt"
	"os"

	"github.com/Ismaelleon/todo-cli/tasks"
	"github.com/spf13/cobra"
)

func init() {
	// Load tasks list before running commands
	cobra.OnInitialize(tasks.Load)
}

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "A todo list app inside your terminal",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("todo-cli: A todo list app inside your terminal\n\nCommands:\nadd: adds a new task\nset: sets the state of a task\ndelete: deletes a task\nlist: lists the tasks\n")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
