package main

import (
	"fmt"
	"os"
	
	"task-manager/command"

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
	if err := RootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
