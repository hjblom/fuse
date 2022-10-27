package commands

import (
	"os"

	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator"
)

func Generate(configPath string) error {
	// Read configuration file
	c, err := config.ReadConfig(configPath, os.ReadFile)
	if err != nil {
		return err
	}

	// Validate config
	err = c.Module.Validate()
	if err != nil {
		return err
	}

	// Generate code
	g := generator.NewGenerator()
	err = g.Generate(c.Module)
	if err != nil {
		return err
	}

	// Write config
	return config.WriteConfig(c, configPath, os.WriteFile)
}
