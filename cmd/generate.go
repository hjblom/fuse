/*
Copyright © 2022 HJ Blom
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hjblom/fuse/internal/commands"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code for the project based on the configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		err := commands.Generate(DefaultConfigPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
