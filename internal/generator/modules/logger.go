package modules

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/hjblom/fuse/internal/common"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

const (
	logrusQualifier  = "github.com/sirupsen/logrus"
	pathQualifier    = "path/filepath"
	runtimeQualifier = "runtime"
	strconvQualifier = "strconv"
	timeQualifier    = "time"
)

var LoggerGenerator = &loggerGenerator{file: util.File}

type loggerGenerator struct {
	file util.FileReadWriter
}

func (g *loggerGenerator) Name() string {
	return "Logger Generator"
}

func (g *loggerGenerator) Description() string {
	return "Generate the logger.go file."
}

func (g *loggerGenerator) Plugins() map[string]string {
	return map[string]string{}
}

// Create configure logging function
//
//	func ConfigureLogging(logLevel string) {
//		logrus.SetFormatter(&logrus.JSONFormatter{
//			TimestampFormat: time.RFC3339Nano,
//			CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
//				return "", filepath.Base(f.File) + ":" + strconv.Itoa(f.Line)
//			},
//		})
//	}
func (g *loggerGenerator) Generate(mod *config.Module) error {
	path := fmt.Sprintf("internal/%s", "logger.go")
	if g.file.Exists(path) {
		return nil
	}

	// Create file
	j := jen.NewFile("internal")

	// Add header
	j.PackageComment(common.Header)

	// Add logurs configure logging function
	g.addLogrusConfigureLogging(j)

	// Write file
	c := fmt.Sprintf("%#v", j)
	err := g.file.WriteFile(path, []byte(c))
	if err != nil {
		return fmt.Errorf("failed to write logger file: %w", err)
	}

	return nil
}

func (g *loggerGenerator) addLogrusConfigureLogging(j *jen.File) {
	j.ImportName(logrusQualifier, "logrus")
	j.Func().Id("ConfigureLogging").Params(jen.Id("logLevel").String()).Block(
		jen.Qual(logrusQualifier, "SetFormatter").Call(
			jen.Op("&").Qual(logrusQualifier, "JSONFormatter").Values(
				jen.Dict{
					jen.Id("TimestampFormat"): jen.Qual(timeQualifier, "RFC3339Nano"),
					jen.Id("CallerPrettyfier"): jen.Func().Params(jen.Id("f").Op("*").Qual(runtimeQualifier, "Frame")).Params(
						jen.Id("function").String(), jen.Id("file").String(),
					).Block(
						jen.Return(jen.Lit(""), jen.Qual(pathQualifier, "Base").Call(jen.Id("f").Dot("File")).Op("+").Lit(":").Op("+").Qual(strconvQualifier, "Itoa").Call(jen.Id("f").Dot("Line"))),
					),
				},
			),
		),
	)
}
