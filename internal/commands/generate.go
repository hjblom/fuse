package commands

import (
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator"
	"github.com/hjblom/fuse/internal/util"
)

func Generate(configPath string) error {
	fi := util.NewFile()

	// Read configuration file
	c, err := config.ReadConfig(configPath, fi)
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
	return config.WriteConfig(c, configPath, fi)
}
