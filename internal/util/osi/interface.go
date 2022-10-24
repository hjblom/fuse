package osi

import "os"

//go:generate mockgen --build_flags=--mod=mod --package=mock --destination=mock/os.go github.com/hjblom/fuse/internal/util/osi Interface

type Interface interface {
	Exists(path string) bool
	ReadFile(name string) ([]byte, error)
	WriteFile(name string, data []byte, perm os.FileMode) error
	MkdirAll(folderPath string, perm os.FileMode) error
}
