package file

import (
	"errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock "github.com/hjblom/fuse/internal/util/file/mock"
)

func TestNewFileIO(t *testing.T) {
	f := NewFileIO()
	assert.NotNil(t, f)
}

func TestFileIO_Read(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mock.NewMockInterface(ctrl)
	f := FileIO{
		readFile:  mock.ReadFile,
		writeFile: mock.WriteFile,
	}

	// Expect
	mock.EXPECT().ReadFile("exist").Return([]byte("exist"), nil).Times(1)
	mock.EXPECT().ReadFile("do-not-exist").Return(nil, errors.New("error")).Times(1)

	// When
	c, err := f.ReadFile("exist")

	// Then
	assert.NotNil(t, c)
	assert.NoError(t, err)

	// When
	c, err = f.ReadFile("do-not-exist")

	// Then
	assert.Nil(t, c)
	assert.Error(t, err)
}

func TestFileIO_Write(t *testing.T) {
	// Given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := mock.NewMockInterface(ctrl)
	f := FileIO{
		readFile:  mock.ReadFile,
		writeFile: mock.WriteFile,
	}

	// Expect
	mock.EXPECT().WriteFile("no-error", []byte("no-error"), os.FileMode(0644)).Return(nil).Times(1)
	mock.EXPECT().WriteFile("error", []byte("error"), os.FileMode(0644)).Return(errors.New("error")).Times(1)

	// When
	err := f.WriteFile("no-error", []byte("no-error"), 0644)
	// Then
	assert.NoError(t, err)

	// When
	err = f.WriteFile("error", []byte("error"), 0644)
	// Then
	assert.Error(t, err)
}
