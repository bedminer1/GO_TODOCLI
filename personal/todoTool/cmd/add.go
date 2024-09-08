/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/bedminer1/personal/todo/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add <task1>...<taskn>",
	Aliases: []string{"a"},
	Short:   "Adding a todo task to the list of todos",
	Args:    cobra.MinimumNArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		tasksFile, err := cmd.Flags().GetString("tasks-file")
		if err != nil {
			return err
		}

		return addTask(os.Stdout, tasksFile, args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addTask(out io.Writer, tasksFile string, args []string) error {
	tl := &todo.TaskList{}
	if err := tl.Load(tasksFile); err != nil {
		return err
	}

	for _, t := range args {
		if err := tl.Add(t); err != nil {
			return err
		}
		fmt.Fprintln(out, "Added Task:", t)
	}

	return tl.Save(tasksFile)
}
