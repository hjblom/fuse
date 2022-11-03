package module

import (
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util/mock"
	"github.com/stretchr/testify/assert"
)

func TestFuseGenerator(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Given
	mfi := mock.NewMockFileInterface(ctrl)
	g := NewFuseGenerator(mfi)
	pkg := &config.Package{
		Name: "client",
		Path: "internal",
	}
	pkg2 := &config.Package{
		Name: "server",
		Path: "internal",
		Requires: []string{
			"internal/client",
		},
		Tags: []string{
			"service",
		},
	}
	mod := &config.Module{
		Path: "test",
		Packages: []*config.Package{
			pkg,
			pkg2,
		},
	}

	// Expect
	mfi.EXPECT().Exists(gomock.Any()).Return(false).Times(1)
	mfi.EXPECT().Write("internal/fuse.go", gomock.Any(), gomock.Any()).DoAndReturn(
		func(path string, data []byte, perm os.FileMode) error {
			// Debug fmt.Println(string(data))
			return nil
		},
	).Times(1)

	// When
	err := g.Generate(mod)

	// Then
	assert.NoError(t, err)
}
