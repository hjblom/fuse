package commands

import (
	"fmt"
	"os"

	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/graph"
)

func Visualize(configPath string) error {
	// Parse config file
	c, err := config.LoadConfigFile(configPath)
	if err != nil {
		fmt.Println("failed to read configuration file: ", err)
	}

	// Config to graph
	g := graph.NewGraph()
	err = g.AddComponents(c.Packages)
	if err != nil {
		fmt.Println("failed validating existing configuration file: ", err)
	}

	// Convert graph to dot
	svg, err := g.ToSVG()
	if err != nil {
		fmt.Println("failed to convert graph to dot: ", err)
	}

	// Write svg to file
	err = os.WriteFile("graph.svg", svg, 0644)
	if err != nil {
		fmt.Println("failed to write svg to file: ", err)
	}

	return nil
}
