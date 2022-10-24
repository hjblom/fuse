package file

import "os"

type Writer func(filename string, data []byte, perm os.FileMode) error
type Reader func(filename string) ([]byte, error)
type MkdirAll func(folderPath string, perm os.FileMode) error

type FileIO struct {
	readFile  Reader
	writeFile Writer
	mkdirAll  MkdirAll
}

func NewFileIO() *FileIO {
	return &FileIO{
		readFile:  os.ReadFile,
		writeFile: os.WriteFile,
		mkdirAll:  os.MkdirAll,
	}
}

func (f *FileIO) Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (f *FileIO) ReadFile(filename string) ([]byte, error) {
	return f.readFile(filename)
}

func (f *FileIO) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return f.writeFile(filename, data, perm)
}

func (f *FileIO) MkdirAll(folderPath string, perm os.FileMode) error {
	return f.mkdirAll(folderPath, perm)
}
