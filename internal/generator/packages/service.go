package packages

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/common"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

var ServiceGenerator = &serviceGenerator{file: util.File}

type serviceGenerator struct {
	file util.FileReadWriter
}

func (g *serviceGenerator) Name() string {
	return "ServiceGenerator"
}

func (g *serviceGenerator) Description() string {
	return "Generate Start() and Stop() methods for a package in a service.go file."
}

func (g *serviceGenerator) Plugins() map[string]string {
	return map[string]string{}
}

func (g *serviceGenerator) Generate(mod *config.Module, pkg *config.Package) error {
	path := fmt.Sprintf("%s/%s", pkg.RelativePath(), "service.go")
	if !pkg.HasTag("service") || g.file.Exists(path) {
		return nil
	}

	// Create file
	j := jen.NewFile(pkg.Name)

	// Add header
	j.PackageComment(common.Header)

	// Add service interface
	// func (p *<pkg.GoAliasName>) Start() error {}
	j.Func().Parens(
		jen.Id(pkg.GoAliasName()).Op("*").Id(pkg.GoStructName()),
	).Id("Start").Params().Error().Block(
		jen.Comment("TODO: Add start logic"),
		jen.Return(jen.Nil()),
	)

	// Newline
	j.Line()

	// func (p *<pkg.GoAliasName>) Stop() error {}
	j.Func().Parens(
		jen.Id(pkg.GoAliasName()).Op("*").Id(pkg.GoStructName()),
	).Id("Stop").Params().Error().Block(
		jen.Comment("TODO: Add stop logic"),
		jen.Return(jen.Nil()),
	)

	// Write file
	content := fmt.Sprintf("%#v", j)
	return g.file.WriteFile(path, []byte(content))
}
