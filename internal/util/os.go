package util

import "os"

type Writer func(filename string, data []byte, perm os.FileMode) error
type Reader func(filename string) ([]byte, error)
type MkdirAll func(folderPath string, perm os.FileMode) error
type IsNotExist func(err error) bool
type Stat func(name string) (os.FileInfo, error)

type OS interface {
	Writer
	Reader
	MkdirAll
	IsNotExist
	Stat
}
