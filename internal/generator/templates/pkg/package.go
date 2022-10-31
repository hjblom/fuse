package pkg

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/config"
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
	pkgName := pkg.GoPackageName()

	// Create file
	j := jen.NewFile(pkg.Name)

	// Add header
	j.PackageComment(DoNotEditHeader)

	/*
		// Struct
		type PackageName struct {
			cfg Config
			Injections
		}
	*/
	j.Type().Id(pkgName).Struct(
		jen.Id("cfg").Id("Config"),
		// TODO: Add injections
	)

	/*
		// Constructor
		func NewPackageName(cfg Config) *PackageName {
			return &PackageName{cfg: cfg}
		}
	*/
	j.Func().Id("New" + pkgName).Params(
		jen.Id("cfg").Id("Config"),
	).Id("Interface").Block(
		jen.Return(jen.Op("&").Id(pkgName).Values(
			jen.Dict{
				jen.Id("cfg"): jen.Id("cfg"),
			},
		)),
	)

	// Write file
	c := fmt.Sprintf("%#v", j)
	err := g.fi.Write(path, []byte(c), 0644)
	if err != nil {
		return fmt.Errorf("failed to write package file: %w", err)
	}

	return nil
}
