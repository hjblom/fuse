/*
Copyright © 2022 HJ Blom
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hjblom/fuse/internal/generator"
	"github.com/hjblom/fuse/internal/util"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code for the project based on the configuration file",
	PreRun: func(cmd *cobra.Command, args []string) {
		// Read configuration
		err := util.File.ReadYamlStruct(PersistentFlagConfigPath, &PersistentConfig)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			os.Exit(1)
		}
		// Validate configuration
		err = PersistentConfig.Module.Validate()
		if err != nil {
			fmt.Println("Error validating config file. Something went wrong:", err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Generate code
		g := generator.NewGenerator()
		err := g.Generate(PersistentConfig.Module)
		if err != nil {
			fmt.Println("Error generating code:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
