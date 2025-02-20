package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DeleteCommand = &cobra.Command{
	Use: "delete",
	Short: "Delete a task",
	Long: `Delete a task from the task list.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
	},
}
