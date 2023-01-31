package generator

import (
	"fmt"

	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

type Generator struct {
	file util.FileReadWriter

	pGens []PackageGenerator
	mGens []ModuleGenerator
}

func NewGenerator() Interface {
	return &Generator{
		file:  util.File,
		pGens: PackageGenerators,
		mGens: ModuleGenerators,
	}
}

func (g *Generator) Generate(mod *config.Module) error {
	// Generate packages
	for _, pkg := range mod.Packages {
		err := g.file.Mkdir(pkg.RelativePath())
		if err != nil {
			return fmt.Errorf("failed to create package directory: %w", err)
		}
		err = g.generatePackages(mod, pkg)
		if err != nil {
			return fmt.Errorf("failed to generate package: %w", err)
		}
	}

	// Ensure internal directory exists
	g.file.Mkdir("internal")
	err := g.generateModules(mod)
	if err != nil {
		return fmt.Errorf("failed to generate module: %w", err)
	}

	return nil
}

func (g *Generator) generateModules(mod *config.Module) error {
	for _, mGen := range g.mGens {
		err := mGen.Generate(mod)
		if err != nil {
			return fmt.Errorf("generator %s failed: %w", mGen.Name(), err)
		}
	}
	return nil
}

func (g *Generator) generatePackages(mod *config.Module, pkg *config.Package) error {
	for _, pGen := range g.pGens {
		err := pGen.Generate(mod, pkg)
		if err != nil {
			return fmt.Errorf("failed to generate package: %w", err)
		}
	}
	return nil
}

// TODO list generators
// func (g *Generator) ListGenerators() []string {
// 	var gens []string
// 	for _, pGen := range g.pGens {
// 		gens = append(gens, pGen.Name())
// 	}
// 	for _, mGen := range g.mGens {
// 		gens = append(gens, mGen.Name())
// 	}
// 	return gens
// }
