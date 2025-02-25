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
	ListCommand.Flags().BoolP("completed", "c", false, "Show completed tasks")
	ListCommand.Flags().BoolP("not-completed", "n", false, "Show not completed tasks")

	ListCommand.Run = func(command *cobra.Command, args []string) {
		completed, _:= command.Flags().GetBool("completed")
		notCompleted, _:= command.Flags().GetBool("not-completed")

		if !completed && !notCompleted {
			completed = true
			notCompleted = true
		}

		if completed {
			pkg.ListTasks(true)
		}

		if notCompleted {
			pkg.ListTasks(false)
		}
	}
}