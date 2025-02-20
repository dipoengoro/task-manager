package command

import (
	"fmt"

	"github.com/spf13/cobra"

	"task-manager"
)

var completeCommand = &cobra.Command{
	Use: "complete",
	Short: "Mark a task as complete",
	Long: `Mark a task as complete.`,
	Run: func(command *cobra.Command, args []string) {
		fmt.Println("complete called")
	},
}

func init() {
	main.RootCommand.AddCommand(completeCommand)
}