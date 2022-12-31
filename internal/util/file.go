package util

import (
	"bytes"
	"os"

	"gopkg.in/yaml.v3"
)

//go:generate mockgen --build_flags=--mod=mod --package=mock --destination=mock/file.go github.com/hjblom/fuse/internal/util FileInterface

const filePermission = 0644
const directoryPermission = 0755

// FileInterface is an interface for file operations
type FileInterface interface {
	Write(path string, data []byte) error
	Read(path string) ([]byte, error)
	Mkdir(folderPath string) error
	Exists(path string) bool
	WriteYamlStruct(path string, data interface{}) error
	ReadYamlStruct(path string, data interface{}) error
}

var File FileInterface = &file{filePermission: filePermission, directoryPermission: directoryPermission}

// file implements FileInterface
type file struct {
	filePermission      os.FileMode
	directoryPermission os.FileMode
}

func (f *file) Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (f *file) Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (f *file) Write(path string, data []byte) error {
	return os.WriteFile(path, data, f.filePermission)
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
	return f.Write(path, b.Bytes())
}

func (f *file) ReadYamlStruct(path string, data interface{}) error {
	b, err := f.Read(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, data)
}
