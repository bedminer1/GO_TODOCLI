/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"os"

	"github.com/bedminer1/personal/todo/todo"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		tasksFile, err := cmd.Flags().GetString("tasks-file")
		if err != nil {
			return nil
		}

		return listAction(io.Writer(os.Stdout), tasksFile)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listAction(out io.Writer, tasksFile string) error {
	tl := &todo.TaskList{}
	if err := tl.Load(tasksFile); err != nil {
		return err
	}

	return tl.List(out)
}
