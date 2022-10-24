package file

import "os"

//go:generate mockgen --build_flags=--mod=mod --package=file --destination=mock/file.go github.com/hjblom/fuse/internal/util/file Interface

type Interface interface {
	Exists(path string) bool
	ReadFile(name string) ([]byte, error)
	WriteFile(name string, data []byte, perm os.FileMode) error
	MkdirAll(folderPath string, perm os.FileMode) error
}
