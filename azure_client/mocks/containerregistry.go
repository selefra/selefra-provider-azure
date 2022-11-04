package mocks

import (
	context "context"
	reflect "reflect"

	containerregistry "github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	containerservice "github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-03-01/containerservice"
	gomock "github.com/golang/mock/gomock"
)

type MockContainerRegistriesClient struct {
	ctrl     *gomock.Controller
	recorder *MockContainerRegistriesClientMockRecorder
}

type MockContainerRegistriesClientMockRecorder struct {
	mock *MockContainerRegistriesClient
}

func NewMockContainerRegistriesClient(ctrl *gomock.Controller) *MockContainerRegistriesClient {
	mock := &MockContainerRegistriesClient{ctrl: ctrl}
	mock.recorder = &MockContainerRegistriesClientMockRecorder{mock}
	return mock
}

func (m *MockContainerRegistriesClient) EXPECT() *MockContainerRegistriesClientMockRecorder {
	return m.recorder
}

func (m *MockContainerRegistriesClient) List(arg0 context.Context) (containerregistry.RegistryListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(containerregistry.RegistryListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockContainerRegistriesClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockContainerRegistriesClient)(nil).List), arg0)
}

type MockContainerReplicationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockContainerReplicationsClientMockRecorder
}

type MockContainerReplicationsClientMockRecorder struct {
	mock *MockContainerReplicationsClient
}

func NewMockContainerReplicationsClient(ctrl *gomock.Controller) *MockContainerReplicationsClient {
	mock := &MockContainerReplicationsClient{ctrl: ctrl}
	mock.recorder = &MockContainerReplicationsClientMockRecorder{mock}
	return mock
}

func (m *MockContainerReplicationsClient) EXPECT() *MockContainerReplicationsClientMockRecorder {
	return m.recorder
}

func (m *MockContainerReplicationsClient) List(arg0 context.Context, arg1, arg2 string) (containerregistry.ReplicationListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(containerregistry.ReplicationListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockContainerReplicationsClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockContainerReplicationsClient)(nil).List), arg0, arg1, arg2)
}

type MockContainerManagedClustersClient struct {
	ctrl     *gomock.Controller
	recorder *MockContainerManagedClustersClientMockRecorder
}

type MockContainerManagedClustersClientMockRecorder struct {
	mock *MockContainerManagedClustersClient
}

func NewMockContainerManagedClustersClient(ctrl *gomock.Controller) *MockContainerManagedClustersClient {
	mock := &MockContainerManagedClustersClient{ctrl: ctrl}
	mock.recorder = &MockContainerManagedClustersClientMockRecorder{mock}
	return mock
}

func (m *MockContainerManagedClustersClient) EXPECT() *MockContainerManagedClustersClientMockRecorder {
	return m.recorder
}

func (m *MockContainerManagedClustersClient) List(arg0 context.Context) (containerservice.ManagedClusterListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(containerservice.ManagedClusterListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockContainerManagedClustersClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockContainerManagedClustersClient)(nil).List), arg0)
}
