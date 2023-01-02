package packages

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/common"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

var ConfigGenerator = &configGenerator{file: util.File}

type configGenerator struct {
	file util.FileReadWriter
}

func (g *configGenerator) Name() string {
	return "Config Generator"
}

func (g *configGenerator) Description() string {
	return "Generate a config.go file for the package."
}

func (g *configGenerator) Plugins() map[string]string {
	return map[string]string{}
}

func (g *configGenerator) Generate(mod *config.Module, pkg *config.Package) error {
	path := fmt.Sprintf("%s/%s", pkg.RelativePath(), "config.go")
	if g.file.Exists(path) && !pkg.HasTag("config") && len(pkg.Config) == 0 {
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
	err := g.file.Write(path, []byte(c))
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
