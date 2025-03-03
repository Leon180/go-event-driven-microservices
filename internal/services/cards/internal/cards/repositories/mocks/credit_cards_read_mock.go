// Code generated by MockGen. DO NOT EDIT.
// Source: credit_cards_read.go
//
// Generated by this command:
//
//	mockgen -source=credit_cards_read.go -destination=mocks/credit_cards_read_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entities "github.com/Leon180/go-event-driven-microservices/internal/services/cards/internal/cards/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockReadCreditCardByMobileNumberAndActiveSwitch is a mock of ReadCreditCardByMobileNumberAndActiveSwitch interface.
type MockReadCreditCardByMobileNumberAndActiveSwitch struct {
	ctrl     *gomock.Controller
	recorder *MockReadCreditCardByMobileNumberAndActiveSwitchMockRecorder
	isgomock struct{}
}

// MockReadCreditCardByMobileNumberAndActiveSwitchMockRecorder is the mock recorder for MockReadCreditCardByMobileNumberAndActiveSwitch.
type MockReadCreditCardByMobileNumberAndActiveSwitchMockRecorder struct {
	mock *MockReadCreditCardByMobileNumberAndActiveSwitch
}

// NewMockReadCreditCardByMobileNumberAndActiveSwitch creates a new mock instance.
func NewMockReadCreditCardByMobileNumberAndActiveSwitch(ctrl *gomock.Controller) *MockReadCreditCardByMobileNumberAndActiveSwitch {
	mock := &MockReadCreditCardByMobileNumberAndActiveSwitch{ctrl: ctrl}
	mock.recorder = &MockReadCreditCardByMobileNumberAndActiveSwitchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadCreditCardByMobileNumberAndActiveSwitch) EXPECT() *MockReadCreditCardByMobileNumberAndActiveSwitchMockRecorder {
	return m.recorder
}

// ReadCreditCardByMobileNumberAndActiveSwitch mocks base method.
func (m *MockReadCreditCardByMobileNumberAndActiveSwitch) ReadCreditCardByMobileNumberAndActiveSwitch(ctx context.Context, mobileNumber string, activeSwitch *bool) (entities.CreditCards, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCreditCardByMobileNumberAndActiveSwitch", ctx, mobileNumber, activeSwitch)
	ret0, _ := ret[0].(entities.CreditCards)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCreditCardByMobileNumberAndActiveSwitch indicates an expected call of ReadCreditCardByMobileNumberAndActiveSwitch.
func (mr *MockReadCreditCardByMobileNumberAndActiveSwitchMockRecorder) ReadCreditCardByMobileNumberAndActiveSwitch(ctx, mobileNumber, activeSwitch any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCreditCardByMobileNumberAndActiveSwitch", reflect.TypeOf((*MockReadCreditCardByMobileNumberAndActiveSwitch)(nil).ReadCreditCardByMobileNumberAndActiveSwitch), ctx, mobileNumber, activeSwitch)
}

// MockReadCreditCard is a mock of ReadCreditCard interface.
type MockReadCreditCard struct {
	ctrl     *gomock.Controller
	recorder *MockReadCreditCardMockRecorder
	isgomock struct{}
}

// MockReadCreditCardMockRecorder is the mock recorder for MockReadCreditCard.
type MockReadCreditCardMockRecorder struct {
	mock *MockReadCreditCard
}

// NewMockReadCreditCard creates a new mock instance.
func NewMockReadCreditCard(ctrl *gomock.Controller) *MockReadCreditCard {
	mock := &MockReadCreditCard{ctrl: ctrl}
	mock.recorder = &MockReadCreditCardMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadCreditCard) EXPECT() *MockReadCreditCardMockRecorder {
	return m.recorder
}

// ReadCreditCard mocks base method.
func (m *MockReadCreditCard) ReadCreditCard(ctx context.Context, id string) (*entities.CreditCard, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadCreditCard", ctx, id)
	ret0, _ := ret[0].(*entities.CreditCard)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadCreditCard indicates an expected call of ReadCreditCard.
func (mr *MockReadCreditCardMockRecorder) ReadCreditCard(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadCreditCard", reflect.TypeOf((*MockReadCreditCard)(nil).ReadCreditCard), ctx, id)
}
