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
		},
	}
}

func (g *Generator) Generate(module string, sorted []config.Package) error {
	for _, component := range sorted {
		fmt.Println("generating component: ", component.Name)

		// Ensure directory exists
		p := fmt.Sprintf("%s/%s", component.Path, component.Name)
		err := g.os.MkdirAll(p, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		// Run generators on directory
		for name, pkg := range g.pkgTemplates {
			fmt.Println("generating file ", name)
			err := pkg.Generate(module, component)
			if err != nil {
				fmt.Printf("failed to generate file: %v", err)
				return err
			}
		}
	}
	return nil
}
