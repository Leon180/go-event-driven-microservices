// Code generated by MockGen. DO NOT EDIT.
// Source: logger.go
//
// Generated by this command:
//
//	mockgen -source=logger.go -destination=./mocks/logger_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	enums "github.com/Leon180/go-event-driven-microservices/internal/pkg/enums"
	context_loggers "github.com/Leon180/go-event-driven-microservices/internal/pkg/utilities/context_loggers"
	gomock "go.uber.org/mock/gomock"
)

// MockContextLogger is a mock of ContextLogger interface.
type MockContextLogger struct {
	ctrl     *gomock.Controller
	recorder *MockContextLoggerMockRecorder
	isgomock struct{}
}

// MockContextLoggerMockRecorder is the mock recorder for MockContextLogger.
type MockContextLoggerMockRecorder struct {
	mock *MockContextLogger
}

// NewMockContextLogger creates a new mock instance.
func NewMockContextLogger(ctrl *gomock.Controller) *MockContextLogger {
	mock := &MockContextLogger{ctrl: ctrl}
	mock.recorder = &MockContextLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContextLogger) EXPECT() *MockContextLoggerMockRecorder {
	return m.recorder
}

// Debug mocks base method.
func (m *MockContextLogger) Debug(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug.
func (mr *MockContextLoggerMockRecorder) Debug(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockContextLogger)(nil).Debug), args...)
}

// Debugf mocks base method.
func (m *MockContextLogger) Debugf(template string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockContextLoggerMockRecorder) Debugf(template any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockContextLogger)(nil).Debugf), varargs...)
}

// Err mocks base method.
func (m *MockContextLogger) Err(msg string, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Err", msg, err)
}

// Err indicates an expected call of Err.
func (mr *MockContextLoggerMockRecorder) Err(msg, err any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockContextLogger)(nil).Err), msg, err)
}

// Error mocks base method.
func (m *MockContextLogger) Error(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockContextLoggerMockRecorder) Error(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockContextLogger)(nil).Error), args...)
}

// Errorf mocks base method.
func (m *MockContextLogger) Errorf(template string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockContextLoggerMockRecorder) Errorf(template any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockContextLogger)(nil).Errorf), varargs...)
}

// Fatal mocks base method.
func (m *MockContextLogger) Fatal(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatal", varargs...)
}

// Fatal indicates an expected call of Fatal.
func (mr *MockContextLoggerMockRecorder) Fatal(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatal", reflect.TypeOf((*MockContextLogger)(nil).Fatal), args...)
}

// Fatalf mocks base method.
func (m *MockContextLogger) Fatalf(template string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf.
func (mr *MockContextLoggerMockRecorder) Fatalf(template any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockContextLogger)(nil).Fatalf), varargs...)
}

// Info mocks base method.
func (m *MockContextLogger) Info(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockContextLoggerMockRecorder) Info(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockContextLogger)(nil).Info), args...)
}

// Infof mocks base method.
func (m *MockContextLogger) Infof(template string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockContextLoggerMockRecorder) Infof(template any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockContextLogger)(nil).Infof), varargs...)
}

// Printf mocks base method.
func (m *MockContextLogger) Printf(template string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Printf", varargs...)
}

// Printf indicates an expected call of Printf.
func (mr *MockContextLoggerMockRecorder) Printf(template any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Printf", reflect.TypeOf((*MockContextLogger)(nil).Printf), varargs...)
}

// Warn mocks base method.
func (m *MockContextLogger) Warn(args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warn", varargs...)
}

// Warn indicates an expected call of Warn.
func (mr *MockContextLoggerMockRecorder) Warn(args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockContextLogger)(nil).Warn), args...)
}

// WarnMsg mocks base method.
func (m *MockContextLogger) WarnMsg(msg string, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WarnMsg", msg, err)
}

// WarnMsg indicates an expected call of WarnMsg.
func (mr *MockContextLoggerMockRecorder) WarnMsg(msg, err any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarnMsg", reflect.TypeOf((*MockContextLogger)(nil).WarnMsg), msg, err)
}

// Warnf mocks base method.
func (m *MockContextLogger) Warnf(template string, args ...any) {
	m.ctrl.T.Helper()
	varargs := []any{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf.
func (mr *MockContextLoggerMockRecorder) Warnf(template any, args ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockContextLogger)(nil).Warnf), varargs...)
}

// WithContextInfo mocks base method.
func (m *MockContextLogger) WithContextInfo(ctx context.Context, keys ...enums.ContextKey) context_loggers.ContextLogger {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range keys {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WithContextInfo", varargs...)
	ret0, _ := ret[0].(context_loggers.ContextLogger)
	return ret0
}

// WithContextInfo indicates an expected call of WithContextInfo.
func (mr *MockContextLoggerMockRecorder) WithContextInfo(ctx any, keys ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, keys...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContextInfo", reflect.TypeOf((*MockContextLogger)(nil).WithContextInfo), varargs...)
}

// clearContextInfo mocks base method.
func (m *MockContextLogger) clearContextInfo() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "clearContextInfo")
}

// clearContextInfo indicates an expected call of clearContextInfo.
func (mr *MockContextLoggerMockRecorder) clearContextInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "clearContextInfo", reflect.TypeOf((*MockContextLogger)(nil).clearContextInfo))
}
