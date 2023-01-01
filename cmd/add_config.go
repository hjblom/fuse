/*
Copyright © 2022 HJ Blom
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Add configuration to a package.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented yet.")
	},
}

func init() {
	addCmd.AddCommand(configCmd)
}
