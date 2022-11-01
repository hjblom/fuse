package pkg

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator/templates/common"
	"github.com/hjblom/fuse/internal/util"
)

const ConfigFileName = "config.go"
const ConfigTag = "config"

type ConfigGenerator struct {
	file util.FileInterface
}

func NewConfigGenerator(file util.FileInterface) Interface {
	return &ConfigGenerator{file: file}
}

func (g *ConfigGenerator) Generate(mod *config.Module, pkg *config.Package) error {
	path := fmt.Sprintf("%s/%s", pkg.RelativePath(), ConfigFileName)
	if g.file.Exists(path) {
		return nil
	}

	// Determine if config should be generated
	if !pkg.HasTag(ConfigTag) {
		return nil
	}

	// Create file
	j := jen.NewFile(pkg.Name)

	// Add header
	j.PackageComment(common.Header)

	// Add config struct
	j.Type().Id("Config").Struct(
		jen.Comment("TODO: Add methods to interface"),
	)

	// Write file
	c := fmt.Sprintf("%#v", j)
	err := g.file.Write(path, []byte(c), 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
