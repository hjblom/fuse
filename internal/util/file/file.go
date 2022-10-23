package file

import "os"

type Writer func(filename string, data []byte, perm os.FileMode) error
type Reader func(filename string) ([]byte, error)

type FileIO struct {
	readFile  Reader
	writeFile Writer
}

func NewFileIO() *FileIO {
	return &FileIO{
		readFile:  os.ReadFile,
		writeFile: os.WriteFile,
	}
}

func (f *FileIO) ReadFile(filename string) ([]byte, error) {
	return f.readFile(filename)
}

func (f *FileIO) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return f.writeFile(filename, data, perm)
}
