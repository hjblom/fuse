package module

import "github.com/hjblom/fuse/internal/config"

type Interface interface {
	Generate(module string, pkg []config.Package) error
}
