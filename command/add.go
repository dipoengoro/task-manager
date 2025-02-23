package command

import (
	"fmt"
	"taskmanager/pkg"

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

func init() {
	AddCommand.Flags().StringP("task", "t", "", "Judul task")
    AddCommand.Flags().StringP("description", "d", "", "Deskripsi task")
    AddCommand.Flags().StringP("deadline", "l", "", "Deadline task (YYYY-MM-DD)")

	AddCommand.Run = func(command *cobra.Command, args []string) {
		task, _:= command.Flags().GetString("task")
		description, _:= command.Flags().GetString( "description")
		deadline, _:= command.Flags().GetString("deadline")

		pkg.AddTask(task, description, deadline)
	}
}