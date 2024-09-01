/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/bedminer1/cobra/pScan/scan"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"l"},
	Short: "List hosts in hosts list",

	RunE: func(cmd *cobra.Command, args []string) error{
		hostsFile, err := cmd.Flags().GetString("hosts-file")
		if err != nil {
			return err
		}
		return listAction(os.Stdout, hostsFile, args)
	},
}

func init() {
	hostsCmd.AddCommand(listCmd)
}

func listAction(out io.Writer, hostsFile string, args []string) error {
	hl := &scan.HostsList{}
	if err := hl.Load(hostsFile); err != nil {
		return err
	}

	for _, h := range hl.Hosts {
		if _, err := fmt.Fprintln(out, h); err != nil {
			return err
		}
	}

	return nil
}