/*
Copyright © 2022 HJ Blom
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator"
	"github.com/hjblom/fuse/internal/util"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate code for the project based on the configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		// Arguments
		configPath := PersistentFlagConfigPath

		// Read configuration file
		c := config.Config{}
		err := util.File.ReadYamlStruct(configPath, &c)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			os.Exit(1)
		}

		// Validate config
		err = c.Module.Validate()
		if err != nil {
			fmt.Println("Error validating config file:", err)
			os.Exit(1)
		}

		// Generate code
		g := generator.NewGenerator()
		err = g.Generate(c.Module)
		if err != nil {
			fmt.Println("Error generating code:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
