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
	ArgPackageName      string
	FlagPackagePath     string
	FlagPackageAlias    string
	FlagPackageRequires []string
	FlagPackageTags     []string
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
		ArgPackageName = args[0]
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Create package
		p := config.NewPackage(ArgPackageName, FlagPackageAlias, FlagPackagePath, FlagPackageRequires, FlagPackageTags)

		// Add package to config
		err := PersistentConfig.Module.AddPackage(p)
		if err != nil {
			fmt.Println("Error adding package to config:", err)
		}
	},
}

func init() {
	addCmd.AddCommand(packageCmd)
	packageCmd.Flags().StringVarP(&FlagPackagePath, "path", "p", ".", "Where the package should be placed")
	packageCmd.Flags().StringVarP(&FlagPackageAlias, "alias", "a", "", "Alias for package instances (defaults to package name)")
	packageCmd.Flags().StringSliceVarP(&FlagPackageRequires, "requires", "r", []string{}, "List of package ids this package depends on")
	packageCmd.Flags().StringSliceVarP(&FlagPackageTags, "tag", "t", []string{}, "List of tags for this package")
}
