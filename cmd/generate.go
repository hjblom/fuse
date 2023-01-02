/*
Copyright © 2022 HJ Blom
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hjblom/fuse/internal/generator"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:    "generate",
	Short:  "Generate code for the project based on the configuration file",
	PreRun: loadAndValidateConfig,
	Run: func(cmd *cobra.Command, args []string) {
		// Generate code
		g := generator.NewGenerator()
		err := g.Generate(RootPersistentConfig.Module)
		if err != nil {
			fmt.Println("Error generating code:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
