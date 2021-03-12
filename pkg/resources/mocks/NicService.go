// Code generated by MockGen. DO NOT EDIT.
// Source: nic.go

// Package mock_resources is a generated GoMock package.
package mock_resources

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	resources "github.com/ionos-cloud/ionosctl/pkg/resources"
)

// MockNicsService is a mock of NicsService interface.
type MockNicsService struct {
	ctrl     *gomock.Controller
	recorder *MockNicsServiceMockRecorder
}

// MockNicsServiceMockRecorder is the mock recorder for MockNicsService.
type MockNicsServiceMockRecorder struct {
	mock *MockNicsService
}

// NewMockNicsService creates a new mock instance.
func NewMockNicsService(ctrl *gomock.Controller) *MockNicsService {
	mock := &MockNicsService{ctrl: ctrl}
	mock.recorder = &MockNicsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNicsService) EXPECT() *MockNicsServiceMockRecorder {
	return m.recorder
}

// AttachToLoadBalancer mocks base method.
func (m *MockNicsService) AttachToLoadBalancer(datacenterId, loadbalancerId, nicId string) (*resources.Nic, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttachToLoadBalancer", datacenterId, loadbalancerId, nicId)
	ret0, _ := ret[0].(*resources.Nic)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AttachToLoadBalancer indicates an expected call of AttachToLoadBalancer.
func (mr *MockNicsServiceMockRecorder) AttachToLoadBalancer(datacenterId, loadbalancerId, nicId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttachToLoadBalancer", reflect.TypeOf((*MockNicsService)(nil).AttachToLoadBalancer), datacenterId, loadbalancerId, nicId)
}

// Create mocks base method.
func (m *MockNicsService) Create(datacenterId, serverId, name string, ips []string, dhcp bool, lan int32) (*resources.Nic, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", datacenterId, serverId, name, ips, dhcp, lan)
	ret0, _ := ret[0].(*resources.Nic)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create.
func (mr *MockNicsServiceMockRecorder) Create(datacenterId, serverId, name, ips, dhcp, lan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNicsService)(nil).Create), datacenterId, serverId, name, ips, dhcp, lan)
}

// Delete mocks base method.
func (m *MockNicsService) Delete(datacenterId, serverId, nicId string) (*resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", datacenterId, serverId, nicId)
	ret0, _ := ret[0].(*resources.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockNicsServiceMockRecorder) Delete(datacenterId, serverId, nicId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNicsService)(nil).Delete), datacenterId, serverId, nicId)
}

// DetachFromLoadBalancer mocks base method.
func (m *MockNicsService) DetachFromLoadBalancer(datacenterId, loadbalancerId, nicId string) (*resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetachFromLoadBalancer", datacenterId, loadbalancerId, nicId)
	ret0, _ := ret[0].(*resources.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetachFromLoadBalancer indicates an expected call of DetachFromLoadBalancer.
func (mr *MockNicsServiceMockRecorder) DetachFromLoadBalancer(datacenterId, loadbalancerId, nicId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetachFromLoadBalancer", reflect.TypeOf((*MockNicsService)(nil).DetachFromLoadBalancer), datacenterId, loadbalancerId, nicId)
}

// Get mocks base method.
func (m *MockNicsService) Get(datacenterId, serverId, nicId string) (*resources.Nic, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", datacenterId, serverId, nicId)
	ret0, _ := ret[0].(*resources.Nic)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockNicsServiceMockRecorder) Get(datacenterId, serverId, nicId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockNicsService)(nil).Get), datacenterId, serverId, nicId)
}

// GetAttachedToLoadBalancer mocks base method.
func (m *MockNicsService) GetAttachedToLoadBalancer(datacenterId, loadbalancerId, nicId string) (*resources.Nic, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttachedToLoadBalancer", datacenterId, loadbalancerId, nicId)
	ret0, _ := ret[0].(*resources.Nic)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAttachedToLoadBalancer indicates an expected call of GetAttachedToLoadBalancer.
func (mr *MockNicsServiceMockRecorder) GetAttachedToLoadBalancer(datacenterId, loadbalancerId, nicId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttachedToLoadBalancer", reflect.TypeOf((*MockNicsService)(nil).GetAttachedToLoadBalancer), datacenterId, loadbalancerId, nicId)
}

// List mocks base method.
func (m *MockNicsService) List(datacenterId, serverId string) (resources.Nics, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", datacenterId, serverId)
	ret0, _ := ret[0].(resources.Nics)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockNicsServiceMockRecorder) List(datacenterId, serverId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNicsService)(nil).List), datacenterId, serverId)
}

// ListAttachedToLoadBalancer mocks base method.
func (m *MockNicsService) ListAttachedToLoadBalancer(datacenterId, loadbalancerId string) (resources.BalancedNics, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAttachedToLoadBalancer", datacenterId, loadbalancerId)
	ret0, _ := ret[0].(resources.BalancedNics)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAttachedToLoadBalancer indicates an expected call of ListAttachedToLoadBalancer.
func (mr *MockNicsServiceMockRecorder) ListAttachedToLoadBalancer(datacenterId, loadbalancerId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAttachedToLoadBalancer", reflect.TypeOf((*MockNicsService)(nil).ListAttachedToLoadBalancer), datacenterId, loadbalancerId)
}

// Update mocks base method.
func (m *MockNicsService) Update(datacenterId, serverId, nicId string, input resources.NicProperties) (*resources.Nic, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", datacenterId, serverId, nicId, input)
	ret0, _ := ret[0].(*resources.Nic)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update.
func (mr *MockNicsServiceMockRecorder) Update(datacenterId, serverId, nicId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNicsService)(nil).Update), datacenterId, serverId, nicId, input)
}
