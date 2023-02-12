package util

import (
	"bytes"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

//go:generate mockgen --build_flags=--mod=mod --package=mock --destination=mock/file.go github.com/hjblom/fuse/internal/util FileReadWriter

const filePermission = 0644
const directoryPermission = 0755

// FileReadWriter is an interface for file operations
type FileReadWriter interface {
	WriteFile(path string, data []byte) error
	ReadFile(path string) ([]byte, error)
	Create(path string) (io.Writer, error)
	Mkdir(folderPath string) error
	Exists(path string) bool
	WriteYamlStruct(path string, data interface{}) error
	ReadYamlStruct(path string, data interface{}) error
}

var File FileReadWriter = &file{filePermission: filePermission, directoryPermission: directoryPermission}

// file implements FileInterface
type file struct {
	filePermission      os.FileMode
	directoryPermission os.FileMode
}

func (f *file) Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (f *file) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (f *file) WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, f.filePermission)
}

func (f *file) Create(path string) (io.Writer, error) {
	return os.Create(path)
}

func (f *file) Mkdir(folderPath string) error {
	return os.MkdirAll(folderPath, f.directoryPermission)
}

func (f *file) WriteYamlStruct(path string, data interface{}) error {
	var b bytes.Buffer
	yamlEncoder := yaml.NewEncoder(&b)
	yamlEncoder.SetIndent(2)
	err := yamlEncoder.Encode(data)
	if err != nil {
		return err
	}
	return f.WriteFile(path, b.Bytes())
}

func (f *file) ReadYamlStruct(path string, data interface{}) error {
	b, err := f.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, data)
}
