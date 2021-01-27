// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_resources is a generated GoMock package.
package mock_resources

import (
	gomock "github.com/golang/mock/gomock"
	resources "github.com/ionos-cloud/ionosctl/pkg/resources"
	reflect "reflect"
)

// MockClientService is a mock of ClientService interface
type MockClientService struct {
	ctrl     *gomock.Controller
	recorder *MockClientServiceMockRecorder
}

// MockClientServiceMockRecorder is the mock recorder for MockClientService
type MockClientServiceMockRecorder struct {
	mock *MockClientService
}

// NewMockClientService creates a new mock instance
func NewMockClientService(ctrl *gomock.Controller) *MockClientService {
	mock := &MockClientService{ctrl: ctrl}
	mock.recorder = &MockClientServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientService) EXPECT() *MockClientServiceMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockClientService) Get() *resources.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(*resources.Client)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockClientServiceMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClientService)(nil).Get))
}

// GetConfig mocks base method
func (m *MockClientService) GetConfig() *resources.ClientConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*resources.ClientConfig)
	return ret0
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockClientServiceMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockClientService)(nil).GetConfig))
}
