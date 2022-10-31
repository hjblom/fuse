package pkg

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hjblom/fuse/internal/config"
	"github.com/hjblom/fuse/internal/util/mock"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Given
	mfi := mock.NewMockFileInterface(ctrl)
	g := NewConfigGenerator(mfi)
	pkg := &config.Package{
		Name: "package",
		Path: "internal",
		Tags: []string{"config"},
	}
	mod := &config.Module{
		Path: "test",
		Packages: []*config.Package{
			pkg,
		},
	}

	// Expect
	mfi.EXPECT().Exists(gomock.Any()).Return(false).Times(1)
	mfi.EXPECT().Write("internal/package/config.go", gomock.Any(), gomock.Any()).Return(nil).Times(1)

	// When
	err := g.Generate(mod, pkg)

	// Then
	assert.NoError(t, err)
}
