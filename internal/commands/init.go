package commands

import (
	"os"

	"github.com/hjblom/fuse/internal/config"
)

func Init(modulePath, configPath string) error {
	c := config.NewConfig(modulePath)
	return config.WriteConfig(c, configPath, os.WriteFile)
}
