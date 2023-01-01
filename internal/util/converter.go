package util

import (
	"bytes"
	"fmt"
)

//go:generate mockgen --build_flags=--mod=mod --package=mock --destination=mock/executor.go github.com/hjblom/fuse/internal/util Converter

type Converter interface {
	ToSvg(dot []byte) ([]byte, error)
}

type converter struct {
	exec Executor
}

var Con Converter = &converter{exec: Exec}

func (c *converter) ToSvg(dot []byte) ([]byte, error) {
	r := bytes.NewReader(dot)
	out, err := c.exec.Execute("dot", WithArgs("-Tsvg"), WithStdIn(r))
	if err != nil {
		return nil, fmt.Errorf("failed to execute \"dot\": %w", err)
	}
	return out, nil
}
