package module

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator/templates/common"
	"github.com/hjblom/fuse/internal/util"
)

const ConfigFile = "config.go"

type ConfigGenerator struct {
	file util.FileInterface
}

func NewConfigGenerator(fi util.FileInterface) Interface {
	return &ConfigGenerator{file: fi}
}

/*
	type Config struct {
		Pkg1 *pkg1.Config
		Pkg2 *pkg2.Config
	}
*/
func (g *ConfigGenerator) Generate(mod *config.Module) error {
	path := fmt.Sprintf("internal/%s", ConfigFile)
	if g.file.Exists(path) {
		return nil
	}

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
	err := g.file.Write(path, []byte(c), 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
