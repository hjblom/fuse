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
		injections = append(injections, jen.Id(reqPkg.Name).Qual(reqPkg.FullPath(mod.Path), "Interface"))
		injectionMap[jen.Id(reqPkg.Name)] = jen.Id(reqPkg.Name)
	}

	/*
		// Struct
		type <PackageName> struct {
			cfg *Config
			Injections
		}
	*/
	j.Type().Id(pkgName).Struct(
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
	j.Func().Id("New" + pkgName).Params(
		injections...,
	).Id("Interface").Block(
		jen.Return(jen.Op("&").Id(pkgName).Values(injectionMap)),
	)

	// Write file
	c := fmt.Sprintf("%#v", j)
	err := g.fi.Write(path, []byte(c), 0644)
	if err != nil {
		return fmt.Errorf("failed to write package file: %w", err)
	}

	return nil
}
