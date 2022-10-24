package osi

import "os"

type Writer func(filename string, data []byte, perm os.FileMode) error
type Reader func(filename string) ([]byte, error)
type MkdirAll func(folderPath string, perm os.FileMode) error
type IsNotExist func(err error) bool
type Stat func(name string) (os.FileInfo, error)

// OS is a wrapper around the os package that can be mocked for testing.
type OS struct {
	readFile   Reader
	writeFile  Writer
	mkdirAll   MkdirAll
	isNotExist IsNotExist
	stat       Stat
}

func NewOS() *OS {
	return &OS{
		readFile:   os.ReadFile,
		writeFile:  os.WriteFile,
		mkdirAll:   os.MkdirAll,
		isNotExist: os.IsNotExist,
		stat:       os.Stat,
	}
}

func (f *OS) Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (f *OS) ReadFile(filename string) ([]byte, error) {
	return f.readFile(filename)
}

func (f *OS) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return f.writeFile(filename, data, perm)
}

func (f *OS) MkdirAll(folderPath string, perm os.FileMode) error {
	return f.mkdirAll(folderPath, perm)
}
