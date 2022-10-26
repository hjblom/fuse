package commands

import (
	"github.com/hjblom/fuse/internal/config"
)

func Add(configPath, packageName, path string, requires, tags []string) error {
	c := config.NewConfig()

	// Read config file
	err := c.Read(configPath)
	if err != nil {
		return err
	}

	// Add package to config file
	pkg := config.NewPackage(packageName, path, requires, tags)
	err = c.AddPackage(pkg)
	if err != nil {
		return err
	}

	// Write config to file
	return c.Write(configPath)
}
