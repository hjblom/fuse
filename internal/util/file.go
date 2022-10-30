package util

import "os"

//go:generate mockgen --build_flags=--mod=mod --package=mock --destination=mock/file.go github.com/hjblom/fuse/internal/util FileInterface

// FileInterface is an interface for file operations
type FileInterface interface {
	Write(filename string, data []byte, perm os.FileMode) error
	Read(filename string) ([]byte, error)
	Mkdir(folderPath string, perm os.FileMode) error
	Exists(path string) bool
}

// File implements FileInterface
type File struct{}

func NewFile() FileInterface {
	return &File{}
}

func (f *File) Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (f *File) Read(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (f *File) Write(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

func (f *File) Mkdir(folderPath string, perm os.FileMode) error {
	return os.MkdirAll(folderPath, perm)
}
