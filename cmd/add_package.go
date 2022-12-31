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

var (
	alias    string
	requires []string
	tags     []string
)

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package [package name]",
	Short: "Add a package to the project.",
	Run: func(cmd *cobra.Command, args []string) {
		// Validation
		if len(args) != 1 {
			fmt.Println("Please provide a package name")
			os.Exit(1)
		}

		// Arguments
		name := args[0]
		configPath := PersistentFlagConfigPath

		// Read config file
		c := config.Config{}
		err := util.File.ReadYamlStruct(configPath, &c)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			os.Exit(1)
		}

		// Create package
		p := config.NewPackage(name, alias, path, requires, tags)

		// Add package to config
		err = c.Module.AddPackage(p)
		if err != nil {
			fmt.Println("Error adding package to config:", err)
		}

		// Write config to file
		err = util.File.WriteYamlStruct(configPath, c)
		if err != nil {
			fmt.Println("Error writing config file:", err)
			os.Exit(1)
		}
	},
}

func init() {
	addCmd.AddCommand(packageCmd)
	packageCmd.Flags().StringVarP(&path, "path", "p", "internal", "Where the package should be placed")
	packageCmd.Flags().StringVarP(&alias, "alias", "a", "", "Alias for package instances (defaults to package name)")
	packageCmd.Flags().StringSliceVarP(&requires, "requires", "r", []string{}, "List of packages this package depends on")
	packageCmd.Flags().StringSliceVarP(&tags, "tag", "t", []string{}, "List of tags for this package")
}
