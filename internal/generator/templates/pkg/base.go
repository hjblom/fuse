package pkg

import "github.com/hjblom/fuse/internal/config"

type Interface interface {
	Generate(mod *config.Module, pkg *config.Package) error
}

// Future improvement

// type Generator struct {
// 	templates map[string]Interface
// }

// func NewGenerator(os osi.Interface) Interface {
// 	return &Generator{
// 		templates: map[string]Interface{
// 			"interface": NewInterfaceGenerator(os),
// 			"config":    NewConfigGenerator(os),
// 			"package":   NewPackageGenerator(os),
// 		},
// 	}
// }

// func (g *Generator) Generate(mod string, sorted []config.Package) error {
// 	for _, pkg := range sorted {
// 		// Ensure directory exists
// 		p := pkg.Path + "/" + pkg.Name
// 		err := osi.MkdirAll(p, 0755)
// 		if err != nil {
// 			return err
// 		}

// 		// Run generators on directory
// 		for _, tpl := range g.templates {
// 			err := tpl.Generate(module, pkg)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }
