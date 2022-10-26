package commands

import (
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator"
)

func Generate(configPath string) error {
	// Read configuration file
	c := config.NewConfig()
	err := c.Read(configPath)
	if err != nil {
		return err
	}

	// Validate config
	err = c.Validate()
	if err != nil {
		return err
	}

	// Generate code
	g := generator.NewGenerator()
	err = g.Generate(c.Module, c.Packages)
	if err != nil {
		return err
	}

	// Write config
	err = c.Write(configPath)
	if err != nil {
		return err
	}

	return nil
}
