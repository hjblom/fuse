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

var (
	AddPersistentFlagTarget string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:              "add",
	Short:            "Add components to the project",
	PersistentPreRun: readAndValidateConfig,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		// Generate
		g := generator.NewGenerator()
		err := g.Generate(RootPersistentConfig.Module)
		if err != nil {
			fmt.Println("Error generating code:", err)
			os.Exit(1)
		}
		// Write configuration
		err = util.File.WriteYamlStruct(RootPersistentFlagConfigPath, &RootPersistentConfig)
		if err != nil {
			fmt.Println("Error writing config file:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
