package pkg

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator/templates/common"
	"github.com/hjblom/fuse/internal/util"
)

type PackageGenerator struct {
	fi util.FileInterface
}

func NewPackageGenerator(fi util.FileInterface) Interface {
	return &PackageGenerator{fi: fi}
}

func (g *PackageGenerator) Generate(mod *config.Module, pkg *config.Package) error {
	path := fmt.Sprintf("%s/%s", pkg.RelativePath(), pkg.GoFileName())
	if g.fi.Exists(path) {
		return nil
	}

	// Create file
	j := jen.NewFile(pkg.Name)

	// Add header
	j.PackageComment(common.DoNotEditHeader)

	// Build injections
	injections := []jen.Code{}
	injectionMap := jen.Dict{}
	// Add optional config injection
	if pkg.HasTag(ConfigTag) {
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
	c := fmt.Sprintf("%#v", j)
	err := g.fi.Write(path, []byte(c), 0644)
	if err != nil {
		return fmt.Errorf("failed to write package file: %w", err)
	}

	return nil
}
