// Code generated by MockGen. DO NOT EDIT.
// Source: uuid.go
//
// Generated by this command:
//
//	mockgen -source=uuid.go -destination=./mocks/uuid_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUUIDGenerator is a mock of UUIDGenerator interface.
type MockUUIDGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockUUIDGeneratorMockRecorder
	isgomock struct{}
}

// MockUUIDGeneratorMockRecorder is the mock recorder for MockUUIDGenerator.
type MockUUIDGeneratorMockRecorder struct {
	mock *MockUUIDGenerator
}

// NewMockUUIDGenerator creates a new mock instance.
func NewMockUUIDGenerator(ctrl *gomock.Controller) *MockUUIDGenerator {
	mock := &MockUUIDGenerator{ctrl: ctrl}
	mock.recorder = &MockUUIDGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUUIDGenerator) EXPECT() *MockUUIDGeneratorMockRecorder {
	return m.recorder
}

// GenerateUUID mocks base method.
func (m *MockUUIDGenerator) GenerateUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GenerateUUID indicates an expected call of GenerateUUID.
func (mr *MockUUIDGeneratorMockRecorder) GenerateUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateUUID", reflect.TypeOf((*MockUUIDGenerator)(nil).GenerateUUID))
}
