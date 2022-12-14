package modules

import (
	"fmt"

	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

var ModTidyGenerator = &modTidyGenerator{file: util.File, exec: util.Exec}

type modTidyGenerator struct {
	file util.FileReadWriter
	exec util.Executor
}

func (g *modTidyGenerator) Name() string {
	return "Mod Tidy"
}

func (g *modTidyGenerator) Description() string {
	return "Executes go mod tidy command."
}

func (g *modTidyGenerator) Plugins() map[string]string {
	return map[string]string{}
}

func (g *modTidyGenerator) Generate(mod *config.Module) error {
	if !g.file.Exists("go.mod") {
		return nil
	}
	_, err := g.exec.Execute("go", util.WithArgs("mod", "tidy"))
	if err != nil {
		return fmt.Errorf("failed to execute \"go mod tidy\": %w", err)
	}
	return nil
}
