// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hjblom/fuse/internal/util (interfaces: FileInterface)

// Package mock is a generated GoMock package.
package mock

import (
	fs "io/fs"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileInterface is a mock of FileInterface interface.
type MockFileInterface struct {
	ctrl     *gomock.Controller
	recorder *MockFileInterfaceMockRecorder
}

// MockFileInterfaceMockRecorder is the mock recorder for MockFileInterface.
type MockFileInterfaceMockRecorder struct {
	mock *MockFileInterface
}

// NewMockFileInterface creates a new mock instance.
func NewMockFileInterface(ctrl *gomock.Controller) *MockFileInterface {
	mock := &MockFileInterface{ctrl: ctrl}
	mock.recorder = &MockFileInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileInterface) EXPECT() *MockFileInterfaceMockRecorder {
	return m.recorder
}

// Exists mocks base method.
func (m *MockFileInterface) Exists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockFileInterfaceMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockFileInterface)(nil).Exists), arg0)
}

// Mkdir mocks base method.
func (m *MockFileInterface) Mkdir(arg0 string, arg1 fs.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mkdir", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mkdir indicates an expected call of Mkdir.
func (mr *MockFileInterfaceMockRecorder) Mkdir(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mkdir", reflect.TypeOf((*MockFileInterface)(nil).Mkdir), arg0, arg1)
}

// Read mocks base method.
func (m *MockFileInterface) Read(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockFileInterfaceMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockFileInterface)(nil).Read), arg0)
}

// Write mocks base method.
func (m *MockFileInterface) Write(arg0 string, arg1 []byte, arg2 fs.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockFileInterfaceMockRecorder) Write(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockFileInterface)(nil).Write), arg0, arg1, arg2)
}