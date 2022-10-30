package generator

import (
	"fmt"

	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator/templates/pkg"
	"github.com/hjblom/fuse/internal/util"
)

type Generator struct {
	fi           util.FileInterface
	pkgTemplates map[string]pkg.Interface
}

func NewGenerator() *Generator {
	fi := util.NewFile()
	return &Generator{
		fi: fi,
		pkgTemplates: map[string]pkg.Interface{
			"interface": pkg.NewInterfaceGenerator(fi),
			"config":    pkg.NewConfigGenerator(fi),
			"package":   pkg.NewPackageGenerator(fi),
		},
	}
}

func (g *Generator) Generate(module *config.Module) error {
	for _, pkg := range module.Packages {
		// Ensure directory exists
		p := fmt.Sprintf("%s/%s", pkg.Path, pkg.Name)
		err := g.fi.Mkdir(p, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		// Run generators on directory
		for _, tpl := range g.pkgTemplates {
			err := tpl.Generate(pkg)
			if err != nil {
				return fmt.Errorf("failed to generate file: %v", err)
			}
		}
	}
	return nil
}
