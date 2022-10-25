package pkg

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util/osi"
)

const ConfigFileName = "config.go"

type ConfigGenerator struct {
	os osi.Interface
}

func NewConfigGenerator(os osi.Interface) Interface {
	return &ConfigGenerator{os: os}
}

func (g *ConfigGenerator) Generate(module string, pkg config.Package) error {
	path := fmt.Sprintf("%s/%s/%s", pkg.Path, pkg.Name, ConfigFileName)
	if g.os.Exists(path) {
		return nil
	}

	// Create file
	j := jen.NewFile(pkg.Name)

	// Add header
	j.PackageComment(Header)

	// Add config struct
	j.Type().Id("Config").Struct(
		jen.Comment("TODO: Add methods to interface"),
	)

	// Write file
	c := fmt.Sprintf("%#v", j)
	err := g.os.WriteFile(path, []byte(c), 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
