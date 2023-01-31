package modules

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/common"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

const goFlagsQualifier = "github.com/jessevdk/go-flags"

var ConfigGenerator = &configGenerator{file: util.File}

type configGenerator struct {
	file util.FileReadWriter
}

func (g *configGenerator) Name() string {
	return "Config"
}

func (g *configGenerator) Description() string {
	return "Generate the module config.go file."
}

func (g *configGenerator) Plugins() map[string]string {
	return map[string]string{}
}

func (g *configGenerator) Generate(mod *config.Module) error {
	// Create file
	j := jen.NewFile("internal")

	// Add header
	j.PackageComment(common.Header)

	// Add config fields
	fields := jen.Statement{}
	fields.Add(generateDefaultLogLevelConfig())

	for _, pkg := range mod.Packages {
		j.ImportName(pkg.FullPath(mod.Path), pkg.Name)
		fields.Add(jen.Id(pkg.GoStructName()).Op("*").Qual(pkg.FullPath(mod.Path), "Config"))
	}

	// Add config struct
	j.Type().Id("Config").Struct(
		fields...,
	)

	// Generate load config function
	j.Add(generateLoadConfigFunction())

	// Write file
	c := fmt.Sprintf("%#v", j)
	path := fmt.Sprintf("internal/%s", "config.go")
	err := g.file.Write(path, []byte(c))
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func generateLoadConfigFunction() *jen.Statement {
	s := &jen.Statement{}
	s.Func().Id("LoadConfig").Params().Params(jen.Op("*").Id("Config"), jen.Id("error")).Block(
		jen.Id("conf").Op(":=").Op("&").Id("Config").Values(),
		jen.Id("parser").Op(":=").Qual(goFlagsQualifier, "NewParser").Call(jen.Id("conf"), jen.Qual(goFlagsQualifier, "Default")),
		jen.If(jen.List(jen.Id("_"), jen.Err()).Op(":=").Id("parser").Dot("Parse").Call(), jen.Err().Op("!=").Nil()).Block(
			jen.Return(jen.Nil(), jen.Err()),
		),
		jen.Return(jen.Id("conf"), jen.Nil()),
	)
	return s
}

func generateDefaultLogLevelConfig() *jen.Statement {
	j := &jen.Statement{}
	j.Id("LogLevel").String().Tag(map[string]string{
		"long":        "log-level",
		"env":         "LOG_LEVEL",
		"description": "The log level to use. Valid values are: DEBUG, INFO, WARN, ERROR",
	})
	return j
}
