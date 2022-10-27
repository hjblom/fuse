package config

import (
	"fmt"
	"strings"
)

type Package struct {
	Name     string   `yaml:"package"`
	Path     string   `yaml:"path,omitempty"`
	Tags     []string `yaml:"tags,omitempty"`
	Requires []string `yaml:"requires,omitempty"`

	module *Module
}

func NewPackage(name, path string, requires, tags []string) *Package {
	return &Package{
		Name:     name,
		Path:     path,
		Requires: requires,
		Tags:     tags,
	}
}

func (p *Package) ModuleName() string {
	return p.module.Path
}

// FullPath returns the full path (including module path) to the package.
func (p *Package) FullPath() string {
	return fmt.Sprintf("%s/%s/%s", p.module.Path, p.Path, p.Name)
}

// RelativePath returns the relative path from the module root to the package.
func (p *Package) RelativePath() string {
	return fmt.Sprintf("%s/%s", p.Path, p.Name)
}

func (p *Package) GoPackageName() string {
	if len(p.Path) == 1 {
		return strings.ToUpper(p.Name)
	}
	return strings.ToUpper(p.Name[0:1]) + string(p.Name[1:])
}

func (p *Package) Validate() error {
	if p.Name == "" {
		return fmt.Errorf("package name is required")
	}
	return nil
}
