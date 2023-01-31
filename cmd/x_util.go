package cmd

import (
	"fmt"
	"os"

	"github.com/hjblom/fuse/internal/util"
	"github.com/spf13/cobra"
)

func readAndValidateConfig(cmd *cobra.Command, args []string) {
	// Read configuration
	err := util.File.ReadYamlStruct(RootPersistentFlagConfigPath, &RootPersistentConfig)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		os.Exit(1)
	}
	// Validate configuration
	err = RootPersistentConfig.Module.Validate()
	if err != nil {
		fmt.Println("Error validating config file. Something went wrong:", err)
		os.Exit(1)
	}
}
