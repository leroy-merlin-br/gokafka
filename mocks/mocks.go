// Code generated by MockGen. DO NOT EDIT.
// Source: consumer/interfaces.go

// Package mock_consumer is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	sarama "github.com/Shopify/sarama"
	gomock "github.com/golang/mock/gomock"
)

// MockConsumerInterface is a mock of ConsumerInterface interface.
type MockConsumerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerInterfaceMockRecorder
}

// MockConsumerInterfaceMockRecorder is the mock recorder for MockConsumerInterface.
type MockConsumerInterfaceMockRecorder struct {
	mock *MockConsumerInterface
}

// NewMockConsumerInterface creates a new mock instance.
func NewMockConsumerInterface(ctrl *gomock.Controller) *MockConsumerInterface {
	mock := &MockConsumerInterface{ctrl: ctrl}
	mock.recorder = &MockConsumerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumerInterface) EXPECT() *MockConsumerInterfaceMockRecorder {
	return m.recorder
}

// Cleanup mocks base method.
func (m *MockConsumerInterface) Cleanup(arg0 sarama.ConsumerGroupSession) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Cleanup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Cleanup indicates an expected call of Cleanup.
func (mr *MockConsumerInterfaceMockRecorder) Cleanup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cleanup", reflect.TypeOf((*MockConsumerInterface)(nil).Cleanup), arg0)
}

// ConsumeClaim mocks base method.
func (m *MockConsumerInterface) ConsumeClaim(arg0 sarama.ConsumerGroupSession, arg1 sarama.ConsumerGroupClaim) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumeClaim", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConsumeClaim indicates an expected call of ConsumeClaim.
func (mr *MockConsumerInterfaceMockRecorder) ConsumeClaim(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeClaim", reflect.TypeOf((*MockConsumerInterface)(nil).ConsumeClaim), arg0, arg1)
}

// IsReady mocks base method.
func (m *MockConsumerInterface) IsReady() chan bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsReady")
	ret0, _ := ret[0].(chan bool)
	return ret0
}

// IsReady indicates an expected call of IsReady.
func (mr *MockConsumerInterfaceMockRecorder) IsReady() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsReady", reflect.TypeOf((*MockConsumerInterface)(nil).IsReady))
}

// SetReady mocks base method.
func (m *MockConsumerInterface) SetReady(ready chan bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReady", ready)
}

// SetReady indicates an expected call of SetReady.
func (mr *MockConsumerInterfaceMockRecorder) SetReady(ready interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReady", reflect.TypeOf((*MockConsumerInterface)(nil).SetReady), ready)
}

// Setup mocks base method.
func (m *MockConsumerInterface) Setup(arg0 sarama.ConsumerGroupSession) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Setup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Setup indicates an expected call of Setup.
func (mr *MockConsumerInterfaceMockRecorder) Setup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockConsumerInterface)(nil).Setup), arg0)
}
