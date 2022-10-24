/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code for the project based on the configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		// Load configuration
		c, err := config.LoadConfigFile(PersistentFlagConfigPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Generate code
		g := generator.NewGenerator()
		err = g.Generate(c.Module, c.Components)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
