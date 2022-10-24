package commands

import (
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/graph"
)

func Add(path, name, grp string, requires, tags []string) error {
	// Load configuration file
	cfg, err := config.LoadConfigFile(path)
	if err != nil {
		return err
	}

	// Load config into graph - this will validate the config
	g := graph.NewGraph()
	err = g.AddComponents(cfg.Components)
	if err != nil {
		return err
	}

	// Add component to graph
	cp := config.Component{
		Name:     name,
		Group:    grp,
		Requires: requires,
		Tags:     tags,
	}
	err = g.AddComponent(cp)
	if err != nil {
		return err
	}

	// Write config to file
	cfg.Components = append(cfg.Components, cp)
	return config.WriteConfigFile(path, cfg)

}
