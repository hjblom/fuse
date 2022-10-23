/*
Copyright © 2022 HJ Blom
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new fuse project",
	Run: func(cmd *cobra.Command, args []string) {
		os.Create(DefaultConfigPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
