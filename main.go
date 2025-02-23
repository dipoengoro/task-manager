package main

import (
	"taskmanager/command"
	"taskmanager/pkg"

	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "taskmanager",
	Short: "A simple task manager CLI application",
	Long:  `taskmanager is a CLI application to manage your tasks.`,
}

func init() {
	RootCommand.AddCommand(command.AddCommand, command.ListCommand, command.DeleteCommand, command.CompleteCommand)
}

func main() {
	pkg.LoadTasks()
	RootCommand.Execute()
}
