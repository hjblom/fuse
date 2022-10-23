/*
Copyright © 2022 HJ Blom
*/
package cmd

import (
	"fmt"

	"github.com/hjblom/fuse/internal/commands"

	"github.com/spf13/cobra"
)

// visualizeCmd represents the visualize command
var visualizeCmd = &cobra.Command{
	Use:   "visualize",
	Short: "Visualize the project dependency graph",
	Run: func(cmd *cobra.Command, args []string) {
		err := commands.Visualize(PersistentFlagConfigPath)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(visualizeCmd)
}
