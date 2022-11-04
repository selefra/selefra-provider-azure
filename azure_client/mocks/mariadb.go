package mocks

import (
	context "context"
	reflect "reflect"

	mariadb "github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2020-01-01/mariadb"
	gomock "github.com/golang/mock/gomock"
)

type MockMariaDBConfigurationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockMariaDBConfigurationsClientMockRecorder
}

type MockMariaDBConfigurationsClientMockRecorder struct {
	mock *MockMariaDBConfigurationsClient
}

func NewMockMariaDBConfigurationsClient(ctrl *gomock.Controller) *MockMariaDBConfigurationsClient {
	mock := &MockMariaDBConfigurationsClient{ctrl: ctrl}
	mock.recorder = &MockMariaDBConfigurationsClientMockRecorder{mock}
	return mock
}

func (m *MockMariaDBConfigurationsClient) EXPECT() *MockMariaDBConfigurationsClientMockRecorder {
	return m.recorder
}

func (m *MockMariaDBConfigurationsClient) ListByServer(arg0 context.Context, arg1, arg2 string) (mariadb.ConfigurationListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(mariadb.ConfigurationListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMariaDBConfigurationsClientMockRecorder) ListByServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByServer", reflect.TypeOf((*MockMariaDBConfigurationsClient)(nil).ListByServer), arg0, arg1, arg2)
}

type MockMariaDBServersClient struct {
	ctrl     *gomock.Controller
	recorder *MockMariaDBServersClientMockRecorder
}

type MockMariaDBServersClientMockRecorder struct {
	mock *MockMariaDBServersClient
}

func NewMockMariaDBServersClient(ctrl *gomock.Controller) *MockMariaDBServersClient {
	mock := &MockMariaDBServersClient{ctrl: ctrl}
	mock.recorder = &MockMariaDBServersClientMockRecorder{mock}
	return mock
}

func (m *MockMariaDBServersClient) EXPECT() *MockMariaDBServersClientMockRecorder {
	return m.recorder
}

func (m *MockMariaDBServersClient) List(arg0 context.Context) (mariadb.ServerListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(mariadb.ServerListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMariaDBServersClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMariaDBServersClient)(nil).List), arg0)
}
