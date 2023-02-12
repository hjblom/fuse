package config

import (
	"fmt"
	"strings"
)

type PackageConfig struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
	Env         string `yaml:"env"`
	Required    bool   `yaml:"required"`
}

type Package struct {
	ID       string          `yaml:"id"`
	Name     string          `yaml:"name"`
	Path     string          `yaml:"path,omitempty"`
	Config   []PackageConfig `yaml:"config,omitempty"`
	Alias    string          `yaml:"alias,omitempty"`
	Tags     []string        `yaml:"tags,omitempty"`
	Requires []string        `yaml:"requires,omitempty"`
}

// FullPath returns the full path (including module path) to the package.
func (p *Package) FullPath(modPath string) string {
	return fmt.Sprintf("%s/%s", modPath, p.RelativePath())
}

// RelativePath returns the relative path from the module root to the package.
func (p *Package) RelativePath() string {
	if p.Path == "" {
		return p.Name
	}
	return fmt.Sprintf("%s/%s", p.Path, p.Name)
}

func (p *Package) GoStructName() string {
	if len(p.Name) == 1 {
		return strings.ToUpper(p.Name)
	}
	return strings.ToUpper(p.Name[0:1]) + string(p.Name[1:])
}

func (p *Package) GoNewStructFuncName() string {
	return fmt.Sprintf("New%s", p.GoStructName())
}

func (p *Package) GoFileName() string {
	return fmt.Sprintf("%s.go", p.Name)
}

func (p *Package) GoAliasName() string {
	if p.Alias != "" {
		return p.Alias
	}
	return p.Name
}

func (p *Package) HasTag(tag string) bool {
	for _, t := range p.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

func (p *Package) AddConfig(cfg PackageConfig) {
	p.Config = append(p.Config, cfg)
}

func (p *Package) Envs() []string {
	var envs []string
	for _, c := range p.Config {
		if c.Env != "" {
			envs = append(envs, c.Env)
		}
	}
	return envs
}
