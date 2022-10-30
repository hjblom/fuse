package generator

import "github.com/hjblom/fuse/internal/config"

//go:generate mockgen --build_flags=--mod=mod --package=mock --destination=mock/generator.go github.com/hjblom/fuse/internal/generator Interface

type Interface interface {
	Generate(module string, sorted []config.Package) error
}
