/*
Copyright © 2022 HJ Blom
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
	"github.com/spf13/cobra"
)

// visualizeCmd represents the visualize command
var visualizeCmd = &cobra.Command{
	Use:   "visualize",
	Short: "Visualize the project dependency graph",
	Run: func(cmd *cobra.Command, args []string) {
		// Arguments
		configPath := PersistentFlagConfigPath

		// Read config file
		c := config.Config{}
		err := util.File.ReadYamlStruct(configPath, &c)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			os.Exit(1)
		}

		svg, err := c.Module.ToSVG()
		if err != nil {
			fmt.Println("failed to generate SVG: ", err)
		}

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
