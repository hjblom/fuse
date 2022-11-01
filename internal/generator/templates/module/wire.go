package module

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator/templates/common"
	"github.com/hjblom/fuse/internal/util"
)

const WireFile = "wire.go"

type WireGenerator struct {
	file util.FileInterface
}

func NewWireGenerator(fi util.FileInterface) Interface {
	return &WireGenerator{file: fi}
}

// Generate the wire.go file.
/*
	func Wire(c *Config) ([]Service, error) {
		services := []Service{}

		// Build according to topological order
		pkg := pkg.NewPackage(c)
		err := pkg.Setup()
		if err != nil {
			return nil, err
		}

		pkg2, err := pkg2.NewPackage(c, pkg)
		if err != nil {
			return nil, err
		}
		pkg2.Setup()

		// Add services
		services = append(services, pkg2)

		return services, nil
	}

*/
func (g *WireGenerator) Generate(mod *config.Module) error {
	path := fmt.Sprintf("internal/%s", WireFile)
	if g.file.Exists(path) {
		return nil
	}

	// Build injection set
	packages, err := mod.TopologicalPackageOrder()
	if err != nil {
		return err
	}

	// Create file
	j := jen.NewFile("internal")

	// Add header
	j.PackageComment(common.Header)

	// Setup content within function
	s := jen.Statement(nil)
	for _, pkg := range packages {
		reqs := []jen.Code{}
		// Fix config injection from being a hard coded string
		if pkg.HasTag("config") {
			reqs = append(reqs, jen.Id("cfg"))
		}
		for _, req := range pkg.Requires {
			reqPkg := mod.GetPackage(req)
			reqs = append(reqs, jen.Id(reqPkg.GoAliasName()))
		}
		// NewPackage
		s.Id(pkg.Name).Op(":=").Qual(pkg.FullPath(mod.Path), pkg.GoNewStructFuncName()).Call(
			reqs...,
		)
		// Setup
		s.Err().Op(":=").Id(pkg.Name).Dot("Setup").Call()
		// Check error
		s.If(jen.Id("err").Op("!=").Nil()).Block(
			jen.Return(jen.Nil(), jen.Id("err")),
		)
		// Optionally add to services if it is a service
		if pkg.HasTag("service") {
			s.Id("services").Op("=").Append(jen.Id("services"), jen.Id(pkg.Name))
		}
	}
	s.Add(jen.Return(jen.Id("services"), jen.Nil()))

	// Inject content setup into function
	j.Func().Id("Wire").Params(jen.Id("c").Op("*").Id("Config")).Params(jen.Index().Id("Service")).Block(
		s...,
	)

	// Write to file
	c := fmt.Sprintf("%#v", j)
	err = g.file.Write(path, []byte(c), 0644)
	if err != nil {
		return fmt.Errorf("failed to write interface file: %w", err)
	}
	return nil
}
