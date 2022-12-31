package generator

import (
	"fmt"

	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

type Generator struct {
	file util.FileInterface

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
		err = g.generatePackage(mod, pkg)
		if err != nil {
			return fmt.Errorf("failed to generate package: %w", err)
		}
	}

	// Ensure internal directory exists
	g.file.Mkdir("internal")
	err := g.generateModule(mod)
	if err != nil {
		return fmt.Errorf("failed to generate module: %w", err)
	}

	return nil
}

func (g *Generator) generateModule(mod *config.Module) error {
	for _, mGen := range g.mGens {
		err := mGen.Generate(mod)
		if err != nil {
			return fmt.Errorf("generator %s failed: %w", mGen.Name(), err)
		}
	}
	return nil
}

func (g *Generator) generatePackage(mod *config.Module, pkg *config.Package) error {
	for _, pGen := range g.pGens {
		err := pGen.Generate(mod, pkg)
		if err != nil {
			return fmt.Errorf("failed to generate package: %w", err)
		}
	}
	return nil
}
