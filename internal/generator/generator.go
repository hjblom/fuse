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
	return &Generator{
		os: osi.NewOS(),
		pkgTemplates: map[string]pkg.Interface{
			"interface": pkg.NewInterfaceGenerator(osi.NewOS()),
		},
	}
}

func (g *Generator) Generate(module string, sorted []config.Package) error {
	for _, component := range sorted {
		// Ensure directory exists
		p := fmt.Sprintf("%s/%s", component.Path, component.Name)
		err := g.os.MkdirAll(p, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		// Run generators on directory
		for _, pkg := range g.pkgTemplates {
			err := pkg.Generate(module, component)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
