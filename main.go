package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCommand = &cobra.Command{
	Use:   "taskmanager",
	Short: "A simple task manager CLI application",
	Long:  `taskmanager is a CLI application to manage your tasks.`,
}

func main() {
	if err := RootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
