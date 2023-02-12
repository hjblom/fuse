package modules

import (
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hjblom/fuse/internal/config"
	util "github.com/hjblom/fuse/internal/util/mock"
	"github.com/stretchr/testify/assert"
)

func Test_Main(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Given
	mockFile := util.NewMockFileReadWriter(ctrl)
	generator := mainGenerator{file: mockFile}
	mod := &config.Module{Path: "github.com/someone/someproject"}

	// Expect
	mockFile.EXPECT().Exists("cmd/main.go").Return(false)
	mockFile.EXPECT().Mkdir("cmd").Return(nil)
	mockFile.EXPECT().Create("cmd/main.go").Return(os.Stderr, nil)

	// When
	err := generator.Generate(mod)

	// Then
	assert.NoError(t, err)
}
