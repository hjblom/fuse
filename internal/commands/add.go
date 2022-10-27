package commands

import (
	"os"

	"github.com/hjblom/fuse/internal/config"
)

func Add(configPath, packageName, packagePath string, requires, tags []string) error {
	// Read config file
	c, err := config.ReadConfig(configPath, os.ReadFile)

	// Add package to config file
	pkg := config.NewPackage(packageName, packagePath, requires, tags)
	err = c.Module.AddPackage(pkg)
	if err != nil {
		return err
	}

	// Write config to file
	return config.WriteConfig(c, configPath, os.WriteFile)
}
