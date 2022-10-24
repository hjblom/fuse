package commands

import (
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/graph"
)

func Add(configPath, pkg, path string, requires, tags []string) error {
	// Load configuration file
	cfg, err := config.LoadConfigFile(configPath)
	if err != nil {
		return err
	}

	// Load config into graph - this will validate the config
	g := graph.NewGraph()
	err = g.AddComponents(cfg.Packages)
	if err != nil {
		return err
	}

	// Add component to graph
	cp := config.Package{
		Name:     pkg,
		Path:     path,
		Requires: requires,
		Tags:     tags,
	}
	err = g.AddComponent(cp)
	if err != nil {
		return err
	}

	// Add component to config
	cfg.Packages = append(cfg.Packages, cp)

	// Write config to file
	return config.WriteConfigFile(configPath, cfg)
}
