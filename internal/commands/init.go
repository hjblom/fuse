package commands

import (
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

func Init(modPath, configPath string) error {
	fi := util.NewFile()
	c := config.NewConfig(modPath)
	return config.WriteConfig(c, configPath, fi)
}
