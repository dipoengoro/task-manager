package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AddCommand = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Long:  `Add a new task to the task list.`,
	Run: func(command *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

