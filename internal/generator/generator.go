package generator

import (
	"fmt"

	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator/templates/module"
	"github.com/hjblom/fuse/internal/generator/templates/pkg"
	"github.com/hjblom/fuse/internal/util"
)

type Generator struct {
	fi           util.FileInterface
	pkgTemplates map[string]pkg.Interface
	modTemplates map[string]module.Interface
}

func NewGenerator() Interface {
	fi := util.NewFile()
	return &Generator{
		fi: fi,
		pkgTemplates: map[string]pkg.Interface{
			"interface": pkg.NewInterfaceGenerator(fi),
			"config":    pkg.NewConfigGenerator(fi),
			"package":   pkg.NewPackageGenerator(fi),
		},
		modTemplates: map[string]module.Interface{
			// "wire":   module.NewWireGenerator(fi),
			"config": module.NewConfigGenerator(fi),
		},
	}
}

func (g *Generator) Generate(mod *config.Module) error {
	for _, pkg := range mod.Packages {
		// Ensure directory exists
		p := fmt.Sprintf("%s/%s", pkg.Path, pkg.Name)
		err := g.fi.Mkdir(p, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		// Run generators on directory
		for _, tpl := range g.pkgTemplates {
			err := tpl.Generate(mod, pkg)
			if err != nil {
				return fmt.Errorf("failed to generate file: %v", err)
			}
		}
	}

	// Run generators on module
	for _, tpl := range g.modTemplates {
		err := tpl.Generate(mod)
		if err != nil {
			return fmt.Errorf("failed to generate file: %v", err)
		}
	}

	return nil
}
