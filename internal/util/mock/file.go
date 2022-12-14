// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hjblom/fuse/internal/util (interfaces: FileReadWriter)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileReadWriter is a mock of FileReadWriter interface.
type MockFileReadWriter struct {
	ctrl     *gomock.Controller
	recorder *MockFileReadWriterMockRecorder
}

// MockFileReadWriterMockRecorder is the mock recorder for MockFileReadWriter.
type MockFileReadWriterMockRecorder struct {
	mock *MockFileReadWriter
}

// NewMockFileReadWriter creates a new mock instance.
func NewMockFileReadWriter(ctrl *gomock.Controller) *MockFileReadWriter {
	mock := &MockFileReadWriter{ctrl: ctrl}
	mock.recorder = &MockFileReadWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileReadWriter) EXPECT() *MockFileReadWriterMockRecorder {
	return m.recorder
}

// Exists mocks base method.
func (m *MockFileReadWriter) Exists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists.
func (mr *MockFileReadWriterMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockFileReadWriter)(nil).Exists), arg0)
}

// Mkdir mocks base method.
func (m *MockFileReadWriter) Mkdir(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mkdir", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mkdir indicates an expected call of Mkdir.
func (mr *MockFileReadWriterMockRecorder) Mkdir(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mkdir", reflect.TypeOf((*MockFileReadWriter)(nil).Mkdir), arg0)
}

// Read mocks base method.
func (m *MockFileReadWriter) Read(arg0 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockFileReadWriterMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockFileReadWriter)(nil).Read), arg0)
}

// ReadYamlStruct mocks base method.
func (m *MockFileReadWriter) ReadYamlStruct(arg0 string, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadYamlStruct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadYamlStruct indicates an expected call of ReadYamlStruct.
func (mr *MockFileReadWriterMockRecorder) ReadYamlStruct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadYamlStruct", reflect.TypeOf((*MockFileReadWriter)(nil).ReadYamlStruct), arg0, arg1)
}

// Write mocks base method.
func (m *MockFileReadWriter) Write(arg0 string, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockFileReadWriterMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockFileReadWriter)(nil).Write), arg0, arg1)
}

// WriteYamlStruct mocks base method.
func (m *MockFileReadWriter) WriteYamlStruct(arg0 string, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteYamlStruct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteYamlStruct indicates an expected call of WriteYamlStruct.
func (mr *MockFileReadWriterMockRecorder) WriteYamlStruct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteYamlStruct", reflect.TypeOf((*MockFileReadWriter)(nil).WriteYamlStruct), arg0, arg1)
}
