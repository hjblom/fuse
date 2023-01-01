package packages

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/common"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

var InterfaceGenerator = &interfaceGenerator{file: util.File}

type interfaceGenerator struct {
	file util.FileReadWriter
}

// Name returns the name of the generator
func (g *interfaceGenerator) Name() string {
	return "Interface generator"
}

// PBase returns true if the generator is a base generator
func (g *interfaceGenerator) Base() bool {
	return true
}

// PDescription returns a description of the package generator
func (g *interfaceGenerator) Description() string {
	return "Generate the interface.go file. This file contains the interface that the package implements."
}

// PPlugins returns a map of plugins - description that the package generator supports
func (g *interfaceGenerator) Plugins() map[string]string {
	return map[string]string{}
}

func (g *interfaceGenerator) Generate(mod *config.Module, pkg *config.Package) error {
	// Skip generation if file already exists
	path := fmt.Sprintf("%s/%s/%s", pkg.Path, pkg.Name, "interface.go")
	if g.file.Exists(path) {
		return nil
	}

	// Create file
	j := jen.NewFile(pkg.Name)

	// Add header
	j.PackageComment(common.Header)

	// Gomock comment
	j.Comment(mockGenComment(mod.Path, pkg.Path, pkg.Name))

	// Add interface
	j.Type().Id("Interface").Interface(
		jen.Comment("TODO: Add methods to interface"),
	)

	// Write file
	c := fmt.Sprintf("%#v", j)
	err := g.file.Write(path, []byte(c))
	if err != nil {
		return fmt.Errorf("failed to write interface file: %w", err)
	}

	return nil
}

func mockGenComment(mod, pkgPath, pkg string) string {
	return "//go:generate mockgen --build_flags=--mod=mod --package=" + pkg + " " +
		"--destination=mock/" + "interface.go" + " " +
		mod + "/" + pkgPath + "/" + pkg + " " +
		"Interface\n"
}
