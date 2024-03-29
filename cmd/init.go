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

var InitArgModulePath string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new fuse project",
	PreRun: func(cmd *cobra.Command, args []string) {
		// Validation
		if len(args) != 1 {
			fmt.Println("Please provide a module name")
			os.Exit(1)
		}
		InitArgModulePath = args[0]
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Init config file
		c := config.NewConfig(config.WithModulePath(InitArgModulePath))
		err := util.File.WriteYamlStruct(RootPersistentFlagConfigPath, c)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
