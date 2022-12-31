package modules

import (
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util"
)

var ModInitGenerator = &modInitGenerator{file: util.File, exec: util.Exec}

type modInitGenerator struct {
	file util.FileInterface
	exec util.Executor
}

func (g *modInitGenerator) Name() string {
	return "Mod Init"
}

func (g *modInitGenerator) Description() string {
	return "Inits go module."
}

func (g *modInitGenerator) Plugins() map[string]string {
	return map[string]string{}
}

func (g *modInitGenerator) Generate(mod *config.Module) error {
	if g.file.Exists("go.mod") {
		return nil
	}
	return g.exec.Command("go", "mod", "init", mod.Path)
}
