package commands

import (
	"fmt"
	"os"

	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

func Visualize(configPath string) error {
	fi := util.NewFile()

	// Parse config file
	c, err := config.ReadConfig(configPath, fi)
	if err != nil {
		fmt.Println("failed to read configuration file: ", err)
	}

	svg, err := c.Module.ToSVG()
	if err != nil {
		fmt.Println("failed to generate SVG: ", err)
	}

	// Write svg to file
	err = os.WriteFile("graph.svg", svg, 0644)
	if err != nil {
		fmt.Println("failed to write svg to file: ", err)
	}

	return nil
}
