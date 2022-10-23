package file

import "os"

//go:generate mockgen --build_flags=--mod=mod --package=file --destination=mock/file.go github.com/hjblom/fuse/internal/util/file Interface

type Interface interface {
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, data []byte, perm os.FileMode) error
}
