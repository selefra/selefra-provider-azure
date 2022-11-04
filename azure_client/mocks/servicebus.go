package mocks

import (
	context "context"
	reflect "reflect"

	servicebus "github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
	gomock "github.com/golang/mock/gomock"
)

type MockServicebusNamespacesClient struct {
	ctrl     *gomock.Controller
	recorder *MockServicebusNamespacesClientMockRecorder
}

type MockServicebusNamespacesClientMockRecorder struct {
	mock *MockServicebusNamespacesClient
}

func NewMockServicebusNamespacesClient(ctrl *gomock.Controller) *MockServicebusNamespacesClient {
	mock := &MockServicebusNamespacesClient{ctrl: ctrl}
	mock.recorder = &MockServicebusNamespacesClientMockRecorder{mock}
	return mock
}

func (m *MockServicebusNamespacesClient) EXPECT() *MockServicebusNamespacesClientMockRecorder {
	return m.recorder
}

func (m *MockServicebusNamespacesClient) List(arg0 context.Context) (servicebus.SBNamespaceListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(servicebus.SBNamespaceListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicebusNamespacesClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockServicebusNamespacesClient)(nil).List), arg0)
}

type MockServicebusTopicsClient struct {
	ctrl     *gomock.Controller
	recorder *MockServicebusTopicsClientMockRecorder
}

type MockServicebusTopicsClientMockRecorder struct {
	mock *MockServicebusTopicsClient
}

func NewMockServicebusTopicsClient(ctrl *gomock.Controller) *MockServicebusTopicsClient {
	mock := &MockServicebusTopicsClient{ctrl: ctrl}
	mock.recorder = &MockServicebusTopicsClientMockRecorder{mock}
	return mock
}

func (m *MockServicebusTopicsClient) EXPECT() *MockServicebusTopicsClientMockRecorder {
	return m.recorder
}

func (m *MockServicebusTopicsClient) ListByNamespace(arg0 context.Context, arg1, arg2 string, arg3, arg4 *int32) (servicebus.SBTopicListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByNamespace", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(servicebus.SBTopicListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicebusTopicsClientMockRecorder) ListByNamespace(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByNamespace", reflect.TypeOf((*MockServicebusTopicsClient)(nil).ListByNamespace), arg0, arg1, arg2, arg3, arg4)
}

type MockServicebusAuthorizationRulesClient struct {
	ctrl     *gomock.Controller
	recorder *MockServicebusAuthorizationRulesClientMockRecorder
}

type MockServicebusAuthorizationRulesClientMockRecorder struct {
	mock *MockServicebusAuthorizationRulesClient
}

func NewMockServicebusAuthorizationRulesClient(ctrl *gomock.Controller) *MockServicebusAuthorizationRulesClient {
	mock := &MockServicebusAuthorizationRulesClient{ctrl: ctrl}
	mock.recorder = &MockServicebusAuthorizationRulesClientMockRecorder{mock}
	return mock
}

func (m *MockServicebusAuthorizationRulesClient) EXPECT() *MockServicebusAuthorizationRulesClientMockRecorder {
	return m.recorder
}

func (m *MockServicebusAuthorizationRulesClient) ListAuthorizationRules(arg0 context.Context, arg1, arg2, arg3 string) (servicebus.SBAuthorizationRuleListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAuthorizationRules", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(servicebus.SBAuthorizationRuleListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicebusAuthorizationRulesClientMockRecorder) ListAuthorizationRules(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAuthorizationRules", reflect.TypeOf((*MockServicebusAuthorizationRulesClient)(nil).ListAuthorizationRules), arg0, arg1, arg2, arg3)
}

type MockServicebusAccessKeysClient struct {
	ctrl     *gomock.Controller
	recorder *MockServicebusAccessKeysClientMockRecorder
}

type MockServicebusAccessKeysClientMockRecorder struct {
	mock *MockServicebusAccessKeysClient
}

func NewMockServicebusAccessKeysClient(ctrl *gomock.Controller) *MockServicebusAccessKeysClient {
	mock := &MockServicebusAccessKeysClient{ctrl: ctrl}
	mock.recorder = &MockServicebusAccessKeysClientMockRecorder{mock}
	return mock
}

func (m *MockServicebusAccessKeysClient) EXPECT() *MockServicebusAccessKeysClientMockRecorder {
	return m.recorder
}

func (m *MockServicebusAccessKeysClient) ListKeys(arg0 context.Context, arg1, arg2, arg3, arg4 string) (servicebus.AccessKeys, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKeys", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(servicebus.AccessKeys)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockServicebusAccessKeysClientMockRecorder) ListKeys(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKeys", reflect.TypeOf((*MockServicebusAccessKeysClient)(nil).ListKeys), arg0, arg1, arg2, arg3, arg4)
}
