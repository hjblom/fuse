package util

import (
	"os"
	"os/exec"
)

//go:generate mockgen --build_flags=--mod=mod --package=mock --destination=mock/executor.go github.com/hjblom/fuse/internal/util Executor

type Executor interface {
	Command(command string, args ...string) error
}

var Exec Executor = &executor{}

type executor struct{}

func (e *executor) Command(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stderr = os.Stdout
	return cmd.Run()
}
