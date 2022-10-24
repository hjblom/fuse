package commands

import "github.com/hjblom/fuse/internal/config"

func Init(module, configPath string) error {
	c := &config.Config{
		Module: module,
	}
	return config.WriteConfigFile(configPath, c)
}
