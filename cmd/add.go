/*
Copyright © 2022 HJ Blom
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hjblom/fuse/internal/commands"

	"github.com/spf13/cobra"
)

var (
	path     string
	requires []string
	tags     []string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [package name]",
	Short: "Add packages to the project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("please provide a package name")
			os.Exit(1)
		}
		err := commands.Add(DefaultConfigPath, args[0], path, requires, tags)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&path, "path", "p", "internal", "Where the package should be placed")
	addCmd.Flags().StringSliceVarP(&requires, "requires", "r", []string{}, "List of packages this package depends on")
	addCmd.Flags().StringSliceVarP(&tags, "tag", "t", []string{}, "List of tags for this package")
}
