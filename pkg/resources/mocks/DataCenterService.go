// Code generated by MockGen. DO NOT EDIT.
// Source: datacenter.go

// Package mock_resources is a generated GoMock package.
package mock_resources

import (
	gomock "github.com/golang/mock/gomock"
	resources "github.com/ionos-cloud/ionosctl/pkg/resources"
	reflect "reflect"
)

// MockDatacentersService is a mock of DatacentersService interface
type MockDatacentersService struct {
	ctrl     *gomock.Controller
	recorder *MockDatacentersServiceMockRecorder
}

// MockDatacentersServiceMockRecorder is the mock recorder for MockDatacentersService
type MockDatacentersServiceMockRecorder struct {
	mock *MockDatacentersService
}

// NewMockDatacentersService creates a new mock instance
func NewMockDatacentersService(ctrl *gomock.Controller) *MockDatacentersService {
	mock := &MockDatacentersService{ctrl: ctrl}
	mock.recorder = &MockDatacentersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatacentersService) EXPECT() *MockDatacentersServiceMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockDatacentersService) List() (resources.Datacenters, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].(resources.Datacenters)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List
func (mr *MockDatacentersServiceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDatacentersService)(nil).List))
}

// Get mocks base method
func (m *MockDatacentersService) Get(datacenterId string) (*resources.Datacenter, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", datacenterId)
	ret0, _ := ret[0].(*resources.Datacenter)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get
func (mr *MockDatacentersServiceMockRecorder) Get(datacenterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDatacentersService)(nil).Get), datacenterId)
}

// Create mocks base method
func (m *MockDatacentersService) Create(name, description, region string) (*resources.Datacenter, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, description, region)
	ret0, _ := ret[0].(*resources.Datacenter)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create
func (mr *MockDatacentersServiceMockRecorder) Create(name, description, region interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDatacentersService)(nil).Create), name, description, region)
}

// Update mocks base method
func (m *MockDatacentersService) Update(datacenterId, name, description string) (*resources.Datacenter, *resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", datacenterId, name, description)
	ret0, _ := ret[0].(*resources.Datacenter)
	ret1, _ := ret[1].(*resources.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Update indicates an expected call of Update
func (mr *MockDatacentersServiceMockRecorder) Update(datacenterId, name, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDatacentersService)(nil).Update), datacenterId, name, description)
}

// Delete mocks base method
func (m *MockDatacentersService) Delete(datacenterId string) (*resources.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", datacenterId)
	ret0, _ := ret[0].(*resources.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockDatacentersServiceMockRecorder) Delete(datacenterId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDatacentersService)(nil).Delete), datacenterId)
}
