package mocks

import (
	context "context"
	reflect "reflect"

	postgresql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql"
	gomock "github.com/golang/mock/gomock"
)

type MockPostgreSQLConfigurationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockPostgreSQLConfigurationsClientMockRecorder
}

type MockPostgreSQLConfigurationsClientMockRecorder struct {
	mock *MockPostgreSQLConfigurationsClient
}

func NewMockPostgreSQLConfigurationsClient(ctrl *gomock.Controller) *MockPostgreSQLConfigurationsClient {
	mock := &MockPostgreSQLConfigurationsClient{ctrl: ctrl}
	mock.recorder = &MockPostgreSQLConfigurationsClientMockRecorder{mock}
	return mock
}

func (m *MockPostgreSQLConfigurationsClient) EXPECT() *MockPostgreSQLConfigurationsClientMockRecorder {
	return m.recorder
}

func (m *MockPostgreSQLConfigurationsClient) ListByServer(arg0 context.Context, arg1, arg2 string) (postgresql.ConfigurationListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(postgresql.ConfigurationListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPostgreSQLConfigurationsClientMockRecorder) ListByServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByServer", reflect.TypeOf((*MockPostgreSQLConfigurationsClient)(nil).ListByServer), arg0, arg1, arg2)
}

type MockPostgreSQLServersClient struct {
	ctrl     *gomock.Controller
	recorder *MockPostgreSQLServersClientMockRecorder
}

type MockPostgreSQLServersClientMockRecorder struct {
	mock *MockPostgreSQLServersClient
}

func NewMockPostgreSQLServersClient(ctrl *gomock.Controller) *MockPostgreSQLServersClient {
	mock := &MockPostgreSQLServersClient{ctrl: ctrl}
	mock.recorder = &MockPostgreSQLServersClientMockRecorder{mock}
	return mock
}

func (m *MockPostgreSQLServersClient) EXPECT() *MockPostgreSQLServersClientMockRecorder {
	return m.recorder
}

func (m *MockPostgreSQLServersClient) List(arg0 context.Context) (postgresql.ServerListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(postgresql.ServerListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPostgreSQLServersClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPostgreSQLServersClient)(nil).List), arg0)
}

type MockPostgreSQLFirewallRulesClient struct {
	ctrl     *gomock.Controller
	recorder *MockPostgreSQLFirewallRulesClientMockRecorder
}

type MockPostgreSQLFirewallRulesClientMockRecorder struct {
	mock *MockPostgreSQLFirewallRulesClient
}

func NewMockPostgreSQLFirewallRulesClient(ctrl *gomock.Controller) *MockPostgreSQLFirewallRulesClient {
	mock := &MockPostgreSQLFirewallRulesClient{ctrl: ctrl}
	mock.recorder = &MockPostgreSQLFirewallRulesClientMockRecorder{mock}
	return mock
}

func (m *MockPostgreSQLFirewallRulesClient) EXPECT() *MockPostgreSQLFirewallRulesClientMockRecorder {
	return m.recorder
}

func (m *MockPostgreSQLFirewallRulesClient) ListByServer(arg0 context.Context, arg1, arg2 string) (postgresql.FirewallRuleListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(postgresql.FirewallRuleListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockPostgreSQLFirewallRulesClientMockRecorder) ListByServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByServer", reflect.TypeOf((*MockPostgreSQLFirewallRulesClient)(nil).ListByServer), arg0, arg1, arg2)
}
