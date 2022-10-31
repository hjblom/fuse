package module

import "github.com/hjblom/fuse/internal/config"

type Interface interface {
	Generate(mod *config.Module) error
}
