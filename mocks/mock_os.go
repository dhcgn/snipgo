// Code generated by MockGen. DO NOT EDIT.
// Source: services/os.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	os "os"
	reflect "reflect"
)

// MockOs is a mock of Os interface
type MockOs struct {
	ctrl     *gomock.Controller
	recorder *MockOsMockRecorder
}

// MockOsMockRecorder is the mock recorder for MockOs
type MockOsMockRecorder struct {
	mock *MockOs
}

// NewMockOs creates a new mock instance
func NewMockOs(ctrl *gomock.Controller) *MockOs {
	mock := &MockOs{ctrl: ctrl}
	mock.recorder = &MockOsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOs) EXPECT() *MockOsMockRecorder {
	return m.recorder
}

// UserHomeDir mocks base method
func (m *MockOs) UserHomeDir() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserHomeDir")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserHomeDir indicates an expected call of UserHomeDir
func (mr *MockOsMockRecorder) UserHomeDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserHomeDir", reflect.TypeOf((*MockOs)(nil).UserHomeDir))
}

// Open mocks base method
func (m *MockOs) Open(name string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", name)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open
func (mr *MockOsMockRecorder) Open(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockOs)(nil).Open), name)
}