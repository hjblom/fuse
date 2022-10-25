package generator

import (
	"fmt"

	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator/templates/pkg"
	"github.com/hjblom/fuse/internal/util/osi"
)

type Generator struct {
	os           osi.Interface
	pkgTemplates map[string]pkg.Interface
}

func NewGenerator() *Generator {
	os := osi.NewOS()
	return &Generator{
		os: os,
		pkgTemplates: map[string]pkg.Interface{
			"interface": pkg.NewInterfaceGenerator(os),
			"config":    pkg.NewConfigGenerator(os),
			"package":   pkg.NewPackageGenerator(os),
		},
	}
}

func (g *Generator) Generate(module string, sorted []config.Package) error {
	for _, pkg := range sorted {
		// Ensure directory exists
		p := fmt.Sprintf("%s/%s", pkg.Path, pkg.Name)
		err := g.os.MkdirAll(p, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		// Run generators on directory
		for _, tpl := range g.pkgTemplates {
			err := tpl.Generate(module, pkg)
			if err != nil {
				return fmt.Errorf("failed to generate file: %v", err)
			}
		}
	}
	return nil
}
