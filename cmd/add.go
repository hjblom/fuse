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

var PersistentConfig *config.Config = config.NewConfig()

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add components to the project",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
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
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Generate
		g := generator.NewGenerator()
		err := g.Generate(PersistentConfig.Module)
		if err != nil {
			fmt.Println("Error generating code:", err)
			os.Exit(1)
		}
		// Write configuration
		err = util.File.WriteYamlStruct(PersistentFlagConfigPath, &PersistentConfig)
		if err != nil {
			fmt.Println("Error writing config file:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
