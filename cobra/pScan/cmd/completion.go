/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate zsh completion for your command",
	Long: `For the tool to show commands available; enter 'source <(./pScan completion)' to use`,

	RunE: func(cmd *cobra.Command, args []string) error {
		return completionAction(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)

}

func completionAction(out io.Writer) error {
	return rootCmd.GenZshCompletion(out)
}
