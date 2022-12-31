package modules

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/common"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

var ConfigGenerator = &configGenerator{file: util.File}

type configGenerator struct {
	file util.FileInterface
}

func (g *configGenerator) Name() string {
	return "Config"
}

func (g *configGenerator) Description() string {
	return "Generate the module config.go file."
}

func (g *configGenerator) Plugins() map[string]string {
	return map[string]string{}
}

func (g *configGenerator) Generate(mod *config.Module) error {
	// Create file
	j := jen.NewFile("internal")

	// Add header
	j.PackageComment(common.Header)

	// Add config fields
	fields := jen.Statement(nil)
	for _, pkg := range mod.Packages {
		fields.Add(jen.Id(pkg.GoAliasName()).Op("*").Qual(pkg.FullPath(mod.Path), "Config"))
	}

	// Add config struct
	j.Type().Id("Config").Struct(
		fields...,
	)

	// Write file
	c := fmt.Sprintf("%#v", j)
	path := fmt.Sprintf("internal/%s", "config.go")
	err := g.file.Write(path, []byte(c))
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
