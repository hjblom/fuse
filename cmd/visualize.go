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
	Use:    "visualize",
	Short:  "Visualize the project dependency graph",
	PreRun: readAndValidateConfig,
	Run: func(cmd *cobra.Command, args []string) {
		// Generate dot
		dot, err := RootPersistentConfig.Module.ToDOT()
		if err != nil {
			fmt.Println("failed to generate dot: ", err)
			os.Exit(1)
		}

		// Convert dot to svg
		svg, err := util.Con.ToSvg(dot)
		if err != nil {
			fmt.Println("failed to convert dot to svg: ", err)
			os.Exit(1)
		}

		// Write svg to file
		err = util.File.WriteFile("graph.svg", svg)
		if err != nil {
			fmt.Println("failed to write svg to file: ", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(visualizeCmd)
}
