// Code generated by MockGen. DO NOT EDIT.
// Source: customers_update.go
//
// Generated by this command:
//
//	mockgen -source=customers_update.go -destination=./mocks/customers_update_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entities "github.com/Leon180/go-event-driven-microservices/internal/services/customers/internal/customers/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockUpdateCustomerByID is a mock of UpdateCustomerByID interface.
type MockUpdateCustomerByID struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateCustomerByIDMockRecorder
	isgomock struct{}
}

// MockUpdateCustomerByIDMockRecorder is the mock recorder for MockUpdateCustomerByID.
type MockUpdateCustomerByIDMockRecorder struct {
	mock *MockUpdateCustomerByID
}

// NewMockUpdateCustomerByID creates a new mock instance.
func NewMockUpdateCustomerByID(ctrl *gomock.Controller) *MockUpdateCustomerByID {
	mock := &MockUpdateCustomerByID{ctrl: ctrl}
	mock.recorder = &MockUpdateCustomerByIDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateCustomerByID) EXPECT() *MockUpdateCustomerByIDMockRecorder {
	return m.recorder
}

// UpdateCustomerByID mocks base method.
func (m *MockUpdateCustomerByID) UpdateCustomerByID(ctx context.Context, update entities.UpdateCustomer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCustomerByID", ctx, update)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCustomerByID indicates an expected call of UpdateCustomerByID.
func (mr *MockUpdateCustomerByIDMockRecorder) UpdateCustomerByID(ctx, update any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCustomerByID", reflect.TypeOf((*MockUpdateCustomerByID)(nil).UpdateCustomerByID), ctx, update)
}
