package command

import (
	"fmt"
	"taskmanager/pkg"

	"github.com/spf13/cobra"
)

var ListCommand = &cobra.Command{
	Use: "list",
	Short: "List all tasks",
	Long: `List all tasks in task list.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init()  {
	ListCommand.Run = func(command *cobra.Command, args []string) {
		pkg.ListTasks()
	}
}