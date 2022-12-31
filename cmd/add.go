/*
Copyright © 2022 HJ Blom
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var (
	path string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add packages to the project",
}

func init() {
	rootCmd.AddCommand(addCmd)
}
