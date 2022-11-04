package mocks

import (
	context "context"
	reflect "reflect"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
	gomock "github.com/golang/mock/gomock"
)

type MockMySQLServersClient struct {
	ctrl     *gomock.Controller
	recorder *MockMySQLServersClientMockRecorder
}

type MockMySQLServersClientMockRecorder struct {
	mock *MockMySQLServersClient
}

func NewMockMySQLServersClient(ctrl *gomock.Controller) *MockMySQLServersClient {
	mock := &MockMySQLServersClient{ctrl: ctrl}
	mock.recorder = &MockMySQLServersClientMockRecorder{mock}
	return mock
}

func (m *MockMySQLServersClient) EXPECT() *MockMySQLServersClientMockRecorder {
	return m.recorder
}

func (m *MockMySQLServersClient) List(arg0 context.Context) (mysql.ServerListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(mysql.ServerListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMySQLServersClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMySQLServersClient)(nil).List), arg0)
}

type MockMySQLConfigurationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockMySQLConfigurationsClientMockRecorder
}

type MockMySQLConfigurationsClientMockRecorder struct {
	mock *MockMySQLConfigurationsClient
}

func NewMockMySQLConfigurationsClient(ctrl *gomock.Controller) *MockMySQLConfigurationsClient {
	mock := &MockMySQLConfigurationsClient{ctrl: ctrl}
	mock.recorder = &MockMySQLConfigurationsClientMockRecorder{mock}
	return mock
}

func (m *MockMySQLConfigurationsClient) EXPECT() *MockMySQLConfigurationsClientMockRecorder {
	return m.recorder
}

func (m *MockMySQLConfigurationsClient) ListByServer(arg0 context.Context, arg1, arg2 string) (mysql.ConfigurationListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(mysql.ConfigurationListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMySQLConfigurationsClientMockRecorder) ListByServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByServer", reflect.TypeOf((*MockMySQLConfigurationsClient)(nil).ListByServer), arg0, arg1, arg2)
}
