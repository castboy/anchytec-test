// Code generated by MockGen. DO NOT EDIT.
// Source: jj.go

// Package mock_mock is a generated GoMock package.
package mock_mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMyOperation is a mock of MyOperation interface
type MockMyOperation struct {
	ctrl     *gomock.Controller
	recorder *MockMyOperationMockRecorder
}

// MockMyOperationMockRecorder is the mock recorder for MockMyOperation
type MockMyOperationMockRecorder struct {
	mock *MockMyOperation
}

// NewMockMyOperation creates a new mock instance
func NewMockMyOperation(ctrl *gomock.Controller) *MockMyOperation {
	mock := &MockMyOperation{ctrl: ctrl}
	mock.recorder = &MockMyOperationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMyOperation) EXPECT() *MockMyOperationMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockMyOperation) Add(a, b int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", a, b)
	ret0, _ := ret[0].(int)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockMyOperationMockRecorder) Add(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockMyOperation)(nil).Add), a, b)
}

// Put mocks base method
func (m *MockMyOperation) Put(a int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", a)
}

// Put indicates an expected call of Put
func (mr *MockMyOperationMockRecorder) Put(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockMyOperation)(nil).Put), a)
}

// RtnErr mocks base method
func (m *MockMyOperation) RtnErr() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RtnErr")
	ret0, _ := ret[0].(error)
	return ret0
}

// RtnErr indicates an expected call of RtnErr
func (mr *MockMyOperationMockRecorder) RtnErr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RtnErr", reflect.TypeOf((*MockMyOperation)(nil).RtnErr))
}

// RtnChan mocks base method
func (m *MockMyOperation) RtnChan() <-chan bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RtnChan")
	ret0, _ := ret[0].(<-chan bool)
	return ret0
}

// RtnChan indicates an expected call of RtnChan
func (mr *MockMyOperationMockRecorder) RtnChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RtnChan", reflect.TypeOf((*MockMyOperation)(nil).RtnChan))
}

// Ptr mocks base method
func (m *MockMyOperation) Ptr(a int, b *int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ptr", a, b)
	ret0, _ := ret[0].(int)
	return ret0
}

// Ptr indicates an expected call of Ptr
func (mr *MockMyOperationMockRecorder) Ptr(a, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ptr", reflect.TypeOf((*MockMyOperation)(nil).Ptr), a, b)
}
