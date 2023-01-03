/*
Copyright © 2022 HJ Blom
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/hjblom/fuse/internal/config"
	"github.com/spf13/cobra"
)

var (
	AddPackageArgName             string
	AddPackageFlagPackagePath     string
	AddPackageFlagPackageAlias    string
	AddPackageFlagPackageRequires []string
	AddPackageFlagPackageTags     []string
)

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package [package name]",
	Short: "Add a package to the project.",
	PreRun: func(cmd *cobra.Command, args []string) {
		// Validation
		if len(args) != 1 {
			fmt.Println("Please provide a package name")
			os.Exit(1)
		}
		AddPackageArgName = args[0]
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Create package
		p := &config.Package{
			Name:     AddPackageArgName,
			Path:     AddPackageFlagPackagePath,
			Alias:    AddPackageFlagPackageAlias,
			Requires: AddPackageFlagPackageRequires,
			Tags:     AddPackageFlagPackageTags,
		}

		// Add package to config
		err := RootPersistentConfig.Module.AddPackage(p)
		if err != nil {
			fmt.Println("Error adding package to config:", err)
		}
	},
}

func init() {
	addCmd.AddCommand(packageCmd)
	packageCmd.Flags().StringVarP(&AddPackageFlagPackagePath, "path", "p", ".", "Where the package should be placed")
	packageCmd.Flags().StringVarP(&AddPackageFlagPackageAlias, "alias", "a", "", "Alias for package instances (default [package name])")
	packageCmd.Flags().StringSliceVarP(&AddPackageFlagPackageRequires, "requires", "r", []string{}, "List of package ids this package depends on")
	packageCmd.Flags().StringSliceVarP(&AddPackageFlagPackageTags, "tag", "t", []string{}, "List of tags for this package")
}
