package util

import (
	"io"
	"os/exec"
)

//go:generate mockgen --build_flags=--mod=mod --package=mock --destination=mock/executor.go github.com/hjblom/fuse/internal/util Executor

type ExecuteOption func(*exec.Cmd)

func WithArgs(args ...string) ExecuteOption {
	return func(cmd *exec.Cmd) {
		cmd.Args = append(cmd.Args, args...)
	}
}

func WithStdIn(r io.Reader) ExecuteOption {
	return func(cmd *exec.Cmd) {
		cmd.Stdin = r
	}
}

func WithStdOut(w io.Writer) ExecuteOption {
	return func(cmd *exec.Cmd) {
		cmd.Stdout = w
	}
}

func WithStdErr(w io.Writer) ExecuteOption {
	return func(cmd *exec.Cmd) {
		cmd.Stderr = w
	}
}

type Executor interface {
	Execute(command string, opts ...ExecuteOption) ([]byte, error)
}

var Exec Executor = &executor{}

type executor struct{}

func (e *executor) Execute(command string, opts ...ExecuteOption) ([]byte, error) {
	cmd := exec.Command(command)
	for _, opt := range opts {
		opt(cmd)
	}
	return cmd.Output()
}
