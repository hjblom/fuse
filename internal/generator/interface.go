package generator

import "github.com/hjblom/fuse/internal/config"

//go:generate mockgen --build_flags=--mod=mod --package=mock --destination=mock/generator.go github.com/hjblom/fuse/internal/generator Interface,ModuleGenerator,PackageGenerator

type Interface interface {
	Generate(mod *config.Module) error
}

type ModuleGenerator interface {
	// Name returns the name of the generator
	Name() string
	// MDescription returns a description of the module generator
	Description() string
	// MPlugins returns a map of plugins - description that the module generator supports
	Plugins() map[string]string
	// MGenerate generates the module
	Generate(mod *config.Module) error
}

type PackageGenerator interface {
	// Name returns the name of the generator
	Name() string
	// Description returns a description of the package generator
	Description() string
	// Plugins returns a map of plugins - description that the package generator supports
	Plugins() map[string]string
	// Generate generates the package
	Generate(mod *config.Module, pkg *config.Package) error
}
