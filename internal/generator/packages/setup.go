package packages

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/common"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

var SetupGenerator = &setupGenerator{file: util.File}

type setupGenerator struct {
	file util.FileReadWriter
}

func (g *setupGenerator) Name() string {
	return "SetupGenerator"
}

func (g *setupGenerator) Description() string {
	return "Generate Setup() method for a package in a setup.go file."
}

func (g *setupGenerator) Plugins() map[string]string {
	return map[string]string{}
}

func (g *setupGenerator) Generate(mod *config.Module, pkg *config.Package) error {
	path := fmt.Sprintf("%s/%s", pkg.RelativePath(), "setup.go")
	if !pkg.HasTag("setup") || g.file.Exists(path) {
		return nil
	}

	// Create file
	j := jen.NewFile(pkg.Name)

	// Add header
	j.PackageComment(common.Header)

	// Add setup interface
	// func (p *<pkg.GoAliasName>) Setup() error {}
	j.Func().Parens(
		jen.Id(pkg.GoAliasName()).Op("*").Id(pkg.GoStructName()),
	).Id("Setup").Params().Error().Block(
		jen.Comment("TODO: Add setup logic"),
		jen.Return(jen.Nil()),
	)

	// Write file
	content := fmt.Sprintf("%#v", j)
	return g.file.WriteFile(path, []byte(content))
}
