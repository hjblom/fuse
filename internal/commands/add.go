package commands

import (
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

func Add(configPath, packageName, packagePath string, requires, tags []string) error {
	fi := util.NewFile()

	// Read config file
	c, err := config.ReadConfig(configPath, fi)

	// Add package to config file
	pkg := config.NewPackage(packageName, packagePath, requires, tags)
	err = c.Module.AddPackage(pkg)
	if err != nil {
		return err
	}

	// Write config to file
	return config.WriteConfig(c, configPath, fi)
}
