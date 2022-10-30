package commands

import (
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

func Init(modulePath, configPath string) error {
	fi := util.NewFile()
	c := config.NewConfig(modulePath)
	return config.WriteConfig(c, configPath, fi)
}
