package module

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator/templates/common"
	"github.com/hjblom/fuse/internal/util"
)

const RuntimeGithub = "github.com/hjblom/fuse/runtime"
const FuseFile = "fuse.go"

type FuseGenerator struct {
	file util.FileInterface
}

func NewFuseGenerator(fi util.FileInterface) Interface {
	return &FuseGenerator{file: fi}
}

// Generate the fuse.go file.
/*
	func Fuse(c *Config) ([]runtime.Service, error) {
		services := []Service{}
		var err error

		// Build according to topological order
		pkg := pkg.NewPackage(c.NewPackage)
		err = pkg.Setup()
		if err != nil {
			return nil, err
		}

		pkg2, err := pkg2.NewPackage2(c.NewPackage2, pkg)
		if err != nil {
			return nil, err
		}
		err = pkg2.Setup()
		if err != nil {
			return nil, err
		}

		// Add services
		services = append(services, pkg2)

		return services, nil
	}

*/
func (g *FuseGenerator) Generate(mod *config.Module) error {
	path := fmt.Sprintf("internal/%s", FuseFile)
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

	// Add []service and err initialization
	s.Add(jen.Var().Id("err").Error())
	j.ImportName(RuntimeGithub, "runtime")
	s.Add(jen.Id("services").Op(":=").Index().Qual(RuntimeGithub, "Service").Values())
	s.Add(jen.Line())

	// Add package initializations
	for _, pkg := range packages {
		// Skip packages if they are not used in dependency injection
		d := mod.GetPackageOutDegree(pkg)
		if d == 0 && len(pkg.Requires) == 0 {
			continue
		}

		// Add package initialization
		reqs := jen.Statement(nil)

		// Optionally add config
		if pkg.HasTag("config") {
			reqs.Add(jen.Id("cfg").Dot(pkg.GoAliasName()))
		}
		for _, req := range pkg.Requires {
			reqPkg := mod.GetPackage(req)
			reqs.Add(jen.Id(reqPkg.GoAliasName()))
		}

		// <PackageAlias>, err := <PackageName>(cfg, injections...)
		j.ImportName(pkg.FullPath(mod.Path), pkg.Name)
		s.Add(jen.Id(pkg.GoAliasName()).Op(":=").Qual(pkg.FullPath(mod.Path), pkg.GoNewStructFuncName()).Call(reqs...))

		// Optionally call setup
		// <PackageAlias>.Setup()
		if pkg.HasTag("setup") {
			// Call setup
			s.Add(jen.Id("err").Op("=").Id(pkg.GoAliasName()).Dot("Setup").Call())
			// Check error
			s.Add(jen.If(jen.Id("err").Op("!=").Nil()).Block(
				jen.Return(jen.Nil(), jen.Id("err")),
			))
		}

		// Optionally add to services if it is a service
		// services = append(services, <PackageAlias>)
		if pkg.HasTag("service") {
			s.Add(jen.Id("services").Op("=").Append(jen.Id("services"), jen.Id(pkg.GoAliasName())))
		}

		// Newline
		s.Add(jen.Line())
	}
	s.Add(jen.Return(jen.Id("services"), jen.Err()))

	// Inject content setup into function
	// func Fuse(c *Config) ([]runtime.Service, error)
	j.Add(jen.Func().Id("Fuse").Params(
		jen.Id("cfg").Op("*").Id("Config"),
	).Params(jen.Index().Qual(RuntimeGithub, "Service"), jen.Error()).Block(
		s...,
	))

	// Write to file
	c := fmt.Sprintf("%#v", j)
	err = g.file.Write(path, []byte(c), 0644)
	if err != nil {
		return fmt.Errorf("failed to write interface file: %w", err)
	}
	return nil
}
