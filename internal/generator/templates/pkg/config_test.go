package pkg

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hjblom/fuse/internal/config"
	os "github.com/hjblom/fuse/internal/util/osi/mock"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Given
	mockOS := os.NewMockInterface(ctrl)
	g := NewConfigGenerator(mockOS)
	c := config.Package{
		Name: "package",
		Path: "internal",
	}

	// Expect
	mockOS.EXPECT().Exists(gomock.Any()).Return(false).Times(1)
	mockOS.EXPECT().WriteFile("internal/package/config.go", gomock.Any(), gomock.Any()).Return(nil).Times(1)

	// When
	err := g.Generate("test", c)

	// Then
	assert.NoError(t, err)
}
