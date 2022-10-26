package commands

import "github.com/hjblom/fuse/internal/config"

func Init(module, configPath string) error {
	c := config.NewConfig()
	c.Module = module
	return c.Write(configPath)
}
