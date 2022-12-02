// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	registry "gitlab.xchch.top/zhangsai/go-101/training/week15/micro_v1/v1/registry"
)

// MockRegistry is a mock of Registry interface.
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry.
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance.
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRegistry) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRegistryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRegistry)(nil).Close))
}

// ListService mocks base method.
func (m *MockRegistry) ListService(ctx context.Context, serviceName string) ([]registry.ServiceInstance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListService", ctx, serviceName)
	ret0, _ := ret[0].([]registry.ServiceInstance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListService indicates an expected call of ListService.
func (mr *MockRegistryMockRecorder) ListService(ctx, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListService", reflect.TypeOf((*MockRegistry)(nil).ListService), ctx, serviceName)
}

// Register mocks base method.
func (m *MockRegistry) Register(ctx context.Context, ins registry.ServiceInstance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, ins)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockRegistryMockRecorder) Register(ctx, ins interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRegistry)(nil).Register), ctx, ins)
}

// Subscribe mocks base method.
func (m *MockRegistry) Subscribe(serviceName string) (<-chan registry.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", serviceName)
	ret0, _ := ret[0].(<-chan registry.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockRegistryMockRecorder) Subscribe(serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockRegistry)(nil).Subscribe), serviceName)
}

// Unregister mocks base method.
func (m *MockRegistry) Unregister(ctx context.Context, ins registry.ServiceInstance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unregister", ctx, ins)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unregister indicates an expected call of Unregister.
func (mr *MockRegistryMockRecorder) Unregister(ctx, ins interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unregister", reflect.TypeOf((*MockRegistry)(nil).Unregister), ctx, ins)
}
