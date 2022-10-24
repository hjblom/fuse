package generator

import (
	"fmt"

	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/generator/templates"
	"github.com/hjblom/fuse/internal/util/file"
)

type Generator struct {
	fi         file.Interface
	generators map[string]func(module string, component config.Component, file file.Interface) error
}

func NewGenerator() *Generator {
	return &Generator{
		fi: file.NewFileIO(),
		generators: map[string]func(module string, component config.Component, file file.Interface) error{
			"interface": templates.GenerateInterface,
		},
	}
}

func (g *Generator) Generate(module string, sorted []config.Component) error {
	fi := file.NewFileIO()
	for _, component := range sorted {
		// Ensure directory exists
		p := fmt.Sprintf("%s/%s", component.Path, component.Package)
		err := fi.MkdirAll(p, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}

		// Run generators on directory
		for _, generate := range g.generators {
			err := generate(module, component, fi)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
