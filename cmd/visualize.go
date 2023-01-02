/*
Copyright © 2022 HJ Blom
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hjblom/fuse/internal/util"
	"github.com/spf13/cobra"
)

// visualizeCmd represents the visualize command
var visualizeCmd = &cobra.Command{
	Use:   "visualize",
	Short: "Visualize the project dependency graph",
	PreRun: func(cmd *cobra.Command, args []string) {
		// Validation
		if len(args) != 0 {
			fmt.Println("No arguments expected")
			os.Exit(1)
		}
		err := util.File.ReadYamlStruct(RootPersistentFlagConfigPath, &RootPersistentConfig)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Generate dot
		dot, err := RootPersistentConfig.Module.ToDOT()
		if err != nil {
			fmt.Println("failed to generate dot: ", err)
			os.Exit(1)
		}

		// Convert dot to svg
		svg, err := util.Con.ToSvg(dot)

		// Write svg to file
		err = util.File.Write("graph.svg", svg)
		if err != nil {
			fmt.Println("failed to write svg to file: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(visualizeCmd)
}
