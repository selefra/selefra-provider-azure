package mocks

import (
	context "context"
	reflect "reflect"

	containerservice "github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-03-01/containerservice"
	gomock "github.com/golang/mock/gomock"
)

type MockManagedClustersClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagedClustersClientMockRecorder
}

type MockManagedClustersClientMockRecorder struct {
	mock *MockManagedClustersClient
}

func NewMockManagedClustersClient(ctrl *gomock.Controller) *MockManagedClustersClient {
	mock := &MockManagedClustersClient{ctrl: ctrl}
	mock.recorder = &MockManagedClustersClientMockRecorder{mock}
	return mock
}

func (m *MockManagedClustersClient) EXPECT() *MockManagedClustersClientMockRecorder {
	return m.recorder
}

func (m *MockManagedClustersClient) List(arg0 context.Context) (containerservice.ManagedClusterListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(containerservice.ManagedClusterListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockManagedClustersClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockManagedClustersClient)(nil).List), arg0)
}
