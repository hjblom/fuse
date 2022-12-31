package packages

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/common"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

var PackageGenerator = &packageGenerator{file: util.File}

type packageGenerator struct {
	file util.FileInterface
}

func (g *packageGenerator) Name() string {
	return "Package generator"
}

func (g *packageGenerator) Description() string {
	return "Generate the package.go file. This file contains the package struct and constructor."
}

func (g *packageGenerator) Plugins() map[string]string {
	return map[string]string{}
}

func (g *packageGenerator) Generate(mod *config.Module, pkg *config.Package) error {
	// Create file
	j := jen.NewFile(pkg.Name)

	// Add header
	j.PackageComment(common.DoNotEditHeader)

	// Build injections
	injections := []jen.Code{}
	injectionMap := jen.Dict{}

	if pkg.HasTag("config") {
		injections = append(injections, jen.Id("cfg").Op("*").Id("Config"))
		injectionMap[jen.Id("cfg")] = jen.Id("cfg")
	}

	// Loop through required struct injections
	for _, req := range pkg.Requires {
		reqPkg := mod.GetPackage(req)
		injections = append(injections, jen.Id(reqPkg.GoAliasName()).Qual(reqPkg.FullPath(mod.Path), "Interface"))
		injectionMap[jen.Id(reqPkg.GoAliasName())] = jen.Id(reqPkg.GoAliasName())
	}

	/*
		// Struct
		type <PackageName> struct {
			cfg *Config
			Injections
		}
	*/
	j.Type().Id(pkg.GoStructName()).Struct(
		injections...,
	)

	/*
		// Constructor
		func New<PackageName>(cfg Config, injections...) *<PackageName> {
			return &PackageName{
				cfg: cfg,
				injections...
			}
		}
	*/
	j.Func().Id(pkg.GoNewStructFuncName()).Params(
		injections...,
	).Op("*").Id(pkg.GoStructName()).Block(
		jen.Return(jen.Op("&").Id(pkg.GoStructName()).Values(injectionMap)),
	)

	// Write file
	path := fmt.Sprintf("%s/%s", pkg.RelativePath(), pkg.GoFileName())
	c := fmt.Sprintf("%#v", j)
	err := g.file.Write(path, []byte(c))
	if err != nil {
		return fmt.Errorf("failed to write package file: %w", err)
	}

	return nil
}
