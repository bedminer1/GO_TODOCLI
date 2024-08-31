/*
Copyright © 2024 bedminer1

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hostsCmd represents the hosts command
var hostsCmd = &cobra.Command{
	Use:   "hosts",
	Short: "Manage the hosts list",
	Long: `Add hosts with the add command
Delete hosts with the delete command
List hosts with the list command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hosts called")
	},
}

func init() {
	rootCmd.AddCommand(hostsCmd)

	
}
