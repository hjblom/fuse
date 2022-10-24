package templates

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hjblom/fuse/internal/config"
	file "github.com/hjblom/fuse/internal/util/file/mock"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Given
	mockFile := file.NewMockInterface(ctrl)
	c := config.Component{
		Package: "package",
		Path:    "internal",
	}

	// Expect
	mockFile.EXPECT().Exists(gomock.Any()).Return(false).Times(1)
	mockFile.EXPECT().WriteFile("internal/package/interface.go", gomock.Any(), gomock.Any()).Return(nil).Times(1)

	// When
	err := GenerateInterface("test", c, mockFile)

	// Then
	assert.NoError(t, err)
}
