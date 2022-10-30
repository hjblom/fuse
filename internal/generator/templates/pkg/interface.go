package pkg

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

const InterfaceFileName = "interface.go"

type InterfaceGenerator struct {
	fi util.FileInterface
}

func NewInterfaceGenerator(fi util.FileInterface) Interface {
	return &InterfaceGenerator{fi: fi}
}

func (g *InterfaceGenerator) Generate(pkg *config.Package) error {
	path := fmt.Sprintf("%s/%s/%s", pkg.Path, pkg.Name, InterfaceFileName)
	if g.fi.Exists(path) {
		return nil
	}

	// Create file
	j := jen.NewFile(pkg.Name)

	// Add header
	j.PackageComment(Header)

	// Gomock comment
	// j.Comment(mockGenComment(pkg.ModuleName(), pkg.Path, pkg.Name))

	// Add interface
	j.Type().Id("Interface").Interface(
		jen.Comment("TODO: Add methods to interface"),
	)

	// Write file
	c := fmt.Sprintf("%#v", j)
	err := g.fi.Write(path, []byte(c), 0644)
	if err != nil {
		return fmt.Errorf("failed to write interface file: %w", err)
	}

	return nil
}

func mockGenComment(module, path, pkg string) string {
	return "//go:generate mockgen --build_flags=--mod=mod --package=" + pkg + " " +
		"--destination=mock/" + InterfaceFileName + " " +
		module + "/" + path + "/" + pkg + " " +
		"Interface\n"
}
