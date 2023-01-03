package packages

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/common"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var ConfigGenerator = &configGenerator{file: util.File}

type configGenerator struct {
	file util.FileReadWriter
}

func (g *configGenerator) Name() string {
	return "Config Generator"
}

func (g *configGenerator) Description() string {
	return "Generate a config.go file for the package."
}

func (g *configGenerator) Plugins() map[string]string {
	return map[string]string{}
}

func (g *configGenerator) Generate(mod *config.Module, pkg *config.Package) error {
	path := fmt.Sprintf("%s/%s", pkg.RelativePath(), "config.go")
	if g.file.Exists(path) && !pkg.HasTag("config") && len(pkg.Config) == 0 {
		return nil
	}

	// Create file
	j := jen.NewFile(pkg.Name)

	// Add header
	j.PackageComment(common.Header)

	// Add config struct
	var fields *jen.Statement
	if len(pkg.Config) == 0 {
		fields = jen.Comment("TODO: Add methods to interface")
	} else {
		fields = generateFields(pkg.Config)
	}
	j.Type().Id("Config").Struct(
		fields,
	)

	// Write file
	c := fmt.Sprintf("%#v", j)
	err := g.file.Write(path, []byte(c))
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func generateFields(fields []config.PackageConfig) *jen.Statement {
	j := &jen.Statement{}
	c := cases.Title(language.English)
	for _, field := range fields {
		name := c.String(field.Name)
		j.Id(name).Id(field.Type).Tag(map[string]string{
			"long":        field.Description,
			"env":         field.Env,
			"description": field.Description,
		})
	}
	return j
}
