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
	AddConfigFlagTarget      string
	AddConfigFlagName        string
	AddConfigFlagType        string
	AddConfigFlagDescription string
	AddConfigFlagEnv         string
	AddConfigFlagRequired    bool
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config [package id]",
	Short: "Add configuration.",
	PreRun: func(cmd *cobra.Command, args []string) {
		// Check if target is set
		if AddConfigFlagTarget == "" {
			fmt.Println("target flag is required")
			cmd.Usage()
			os.Exit(1)
		}

		// Check if name is set
		if AddConfigFlagName == "" {
			fmt.Println("name flag is required")
			cmd.Usage()
			os.Exit(1)
		}

		// Check if type is set
		if AddConfigFlagType == "" {
			fmt.Println("type flag is required")
			cmd.Usage()
			os.Exit(1)
		}

		// Check if description is set
		if AddConfigFlagDescription == "" {
			fmt.Println("description flag is required")
			cmd.Usage()
			os.Exit(1)
		}

		// Check if env is set
		if AddConfigFlagEnv == "" {
			fmt.Println("env flag is required")
			cmd.Usage()
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure that the environment variable does not already exist
		// envs := RootPersistentConfig.Module.GetEnvs()
		// if envs[AddConfigFlagEnv] {
		// 	fmt.Println("environment variable already exists, please choose a different value:", AddConfigFlagEnv)
		// 	os.Exit(1)
		// }
		p := RootPersistentConfig.Module.GetPackage(AddConfigFlagTarget)
		if p == nil {
			fmt.Println("target package not found:", AddConfigFlagTarget)
			os.Exit(1)
		}
		c := config.PackageConfig{
			Name:        AddConfigFlagName,
			Type:        AddConfigFlagType,
			Description: AddConfigFlagDescription,
			Env:         AddConfigFlagEnv,
			Required:    AddConfigFlagRequired,
		}
		p.AddConfig(c)
	},
}

func init() {
	addCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&AddConfigFlagTarget, "target", "t", "", "The target package to add configuration to (required)")
	configCmd.Flags().StringVarP(&AddConfigFlagName, "name", "n", "", "The field name of the configuration (required)")
	configCmd.Flags().StringVarP(&AddConfigFlagType, "type", "y", "", "The field type of the configuration (required)")
	configCmd.Flags().StringVarP(&AddConfigFlagDescription, "description", "d", "", "The field description of the configuration (required)")
	configCmd.Flags().StringVarP(&AddConfigFlagEnv, "env", "e", "", "The field environment variable of the configuration (required)")
	configCmd.Flags().BoolVarP(&AddConfigFlagRequired, "required", "r", false, "A field indicating if the configuration is required")
	configCmd.Flags().SortFlags = false
}
