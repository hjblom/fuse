package generator

import (
	"fmt"

	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator/templates"
	"github.com/hjblom/fuse/internal/util/osi"
)

type Generator struct {
	os        osi.Interface
	templates map[string]templates.Interface
}

func NewGenerator() *Generator {
	return &Generator{
		os: osi.NewOS(),
		templates: map[string]templates.Interface{
			"interface": templates.NewInterfaceGenerator(osi.NewOS()),
		},
	}
}

func (g *Generator) Generate(module string, sorted []config.Component) error {
	for _, component := range sorted {
		// Ensure directory exists
		p := fmt.Sprintf("%s/%s", component.Path, component.Package)
		err := g.os.MkdirAll(p, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		// Run generators on directory
		for _, tpl := range g.templates {
			err := tpl.Generate(module, component)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
