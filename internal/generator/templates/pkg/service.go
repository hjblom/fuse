package pkg

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator/templates/common"
	"github.com/hjblom/fuse/internal/util"
)

const SetupFileName = "service.go"
const SetupTag = "service"

type ServiceGenerator struct {
	file util.FileInterface
}

func NewServiceGenerator(file util.FileInterface) Interface {
	return &ServiceGenerator{file: file}
}

func (g *ServiceGenerator) Generate(mod *config.Module, pkg *config.Package) error {
	path := fmt.Sprintf("%s/%s", pkg.RelativePath(), SetupFileName)
	if g.file.Exists(path) {
		return nil
	}

	// Determine if Setup should be generated
	if !pkg.HasTag(SetupTag) {
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
	)

	// Newline
	j.Line()

	// func (p *<pkg.GoAliasName>) Stop() error {}
	j.Func().Parens(
		jen.Id(pkg.GoAliasName()).Op("*").Id(pkg.GoStructName()),
	).Id("Stop").Params().Error().Block(
		jen.Comment("TODO: Add stop logic"),
	)

	// Write file
	c := fmt.Sprintf("%#v", j)
	err := g.file.Write(path, []byte(c), 0644)
	if err != nil {
		return fmt.Errorf("failed to write Setup file: %w", err)
	}

	return nil
}
