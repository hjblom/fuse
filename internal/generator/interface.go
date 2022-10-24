package generator

import "github.com/hjblom/fuse/internal/config"

type Interface interface {
	Generate(module string, sorted []config.Component) error
}
