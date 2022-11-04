package mocks

import (
	context "context"
	reflect "reflect"

	cdn "github.com/Azure/azure-sdk-for-go/services/cdn/mgmt/2020-09-01/cdn"
	gomock "github.com/golang/mock/gomock"
)

type MockCDNProfilesClient struct {
	ctrl     *gomock.Controller
	recorder *MockCDNProfilesClientMockRecorder
}

type MockCDNProfilesClientMockRecorder struct {
	mock *MockCDNProfilesClient
}

func NewMockCDNProfilesClient(ctrl *gomock.Controller) *MockCDNProfilesClient {
	mock := &MockCDNProfilesClient{ctrl: ctrl}
	mock.recorder = &MockCDNProfilesClientMockRecorder{mock}
	return mock
}

func (m *MockCDNProfilesClient) EXPECT() *MockCDNProfilesClientMockRecorder {
	return m.recorder
}

func (m *MockCDNProfilesClient) List(arg0 context.Context) (cdn.ProfileListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(cdn.ProfileListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCDNProfilesClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCDNProfilesClient)(nil).List), arg0)
}

type MockCDNEndpointsClient struct {
	ctrl     *gomock.Controller
	recorder *MockCDNEndpointsClientMockRecorder
}

type MockCDNEndpointsClientMockRecorder struct {
	mock *MockCDNEndpointsClient
}

func NewMockCDNEndpointsClient(ctrl *gomock.Controller) *MockCDNEndpointsClient {
	mock := &MockCDNEndpointsClient{ctrl: ctrl}
	mock.recorder = &MockCDNEndpointsClientMockRecorder{mock}
	return mock
}

func (m *MockCDNEndpointsClient) EXPECT() *MockCDNEndpointsClientMockRecorder {
	return m.recorder
}

func (m *MockCDNEndpointsClient) ListByProfile(arg0 context.Context, arg1, arg2 string) (cdn.EndpointListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByProfile", arg0, arg1, arg2)
	ret0, _ := ret[0].(cdn.EndpointListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCDNEndpointsClientMockRecorder) ListByProfile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByProfile", reflect.TypeOf((*MockCDNEndpointsClient)(nil).ListByProfile), arg0, arg1, arg2)
}

type MockCDNCustomDomainsClient struct {
	ctrl     *gomock.Controller
	recorder *MockCDNCustomDomainsClientMockRecorder
}

type MockCDNCustomDomainsClientMockRecorder struct {
	mock *MockCDNCustomDomainsClient
}

func NewMockCDNCustomDomainsClient(ctrl *gomock.Controller) *MockCDNCustomDomainsClient {
	mock := &MockCDNCustomDomainsClient{ctrl: ctrl}
	mock.recorder = &MockCDNCustomDomainsClientMockRecorder{mock}
	return mock
}

func (m *MockCDNCustomDomainsClient) EXPECT() *MockCDNCustomDomainsClientMockRecorder {
	return m.recorder
}

func (m *MockCDNCustomDomainsClient) ListByEndpoint(arg0 context.Context, arg1, arg2, arg3 string) (cdn.CustomDomainListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cdn.CustomDomainListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCDNCustomDomainsClientMockRecorder) ListByEndpoint(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByEndpoint", reflect.TypeOf((*MockCDNCustomDomainsClient)(nil).ListByEndpoint), arg0, arg1, arg2, arg3)
}

type MockCDNOriginsClient struct {
	ctrl     *gomock.Controller
	recorder *MockCDNOriginsClientMockRecorder
}

type MockCDNOriginsClientMockRecorder struct {
	mock *MockCDNOriginsClient
}

func NewMockCDNOriginsClient(ctrl *gomock.Controller) *MockCDNOriginsClient {
	mock := &MockCDNOriginsClient{ctrl: ctrl}
	mock.recorder = &MockCDNOriginsClientMockRecorder{mock}
	return mock
}

func (m *MockCDNOriginsClient) EXPECT() *MockCDNOriginsClientMockRecorder {
	return m.recorder
}

func (m *MockCDNOriginsClient) ListByEndpoint(arg0 context.Context, arg1, arg2, arg3 string) (cdn.OriginListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cdn.OriginListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCDNOriginsClientMockRecorder) ListByEndpoint(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByEndpoint", reflect.TypeOf((*MockCDNOriginsClient)(nil).ListByEndpoint), arg0, arg1, arg2, arg3)
}

type MockCDNOriginGroupsClient struct {
	ctrl     *gomock.Controller
	recorder *MockCDNOriginGroupsClientMockRecorder
}

type MockCDNOriginGroupsClientMockRecorder struct {
	mock *MockCDNOriginGroupsClient
}

func NewMockCDNOriginGroupsClient(ctrl *gomock.Controller) *MockCDNOriginGroupsClient {
	mock := &MockCDNOriginGroupsClient{ctrl: ctrl}
	mock.recorder = &MockCDNOriginGroupsClientMockRecorder{mock}
	return mock
}

func (m *MockCDNOriginGroupsClient) EXPECT() *MockCDNOriginGroupsClientMockRecorder {
	return m.recorder
}

func (m *MockCDNOriginGroupsClient) ListByEndpoint(arg0 context.Context, arg1, arg2, arg3 string) (cdn.OriginGroupListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cdn.OriginGroupListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCDNOriginGroupsClientMockRecorder) ListByEndpoint(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByEndpoint", reflect.TypeOf((*MockCDNOriginGroupsClient)(nil).ListByEndpoint), arg0, arg1, arg2, arg3)
}

type MockCDNRoutesClient struct {
	ctrl     *gomock.Controller
	recorder *MockCDNRoutesClientMockRecorder
}

type MockCDNRoutesClientMockRecorder struct {
	mock *MockCDNRoutesClient
}

func NewMockCDNRoutesClient(ctrl *gomock.Controller) *MockCDNRoutesClient {
	mock := &MockCDNRoutesClient{ctrl: ctrl}
	mock.recorder = &MockCDNRoutesClientMockRecorder{mock}
	return mock
}

func (m *MockCDNRoutesClient) EXPECT() *MockCDNRoutesClientMockRecorder {
	return m.recorder
}

func (m *MockCDNRoutesClient) ListByEndpoint(arg0 context.Context, arg1, arg2, arg3 string) (cdn.RouteListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cdn.RouteListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCDNRoutesClientMockRecorder) ListByEndpoint(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByEndpoint", reflect.TypeOf((*MockCDNRoutesClient)(nil).ListByEndpoint), arg0, arg1, arg2, arg3)
}

type MockCDNRuleSetsClient struct {
	ctrl     *gomock.Controller
	recorder *MockCDNRuleSetsClientMockRecorder
}

type MockCDNRuleSetsClientMockRecorder struct {
	mock *MockCDNRuleSetsClient
}

func NewMockCDNRuleSetsClient(ctrl *gomock.Controller) *MockCDNRuleSetsClient {
	mock := &MockCDNRuleSetsClient{ctrl: ctrl}
	mock.recorder = &MockCDNRuleSetsClientMockRecorder{mock}
	return mock
}

func (m *MockCDNRuleSetsClient) EXPECT() *MockCDNRuleSetsClientMockRecorder {
	return m.recorder
}

func (m *MockCDNRuleSetsClient) ListByProfile(arg0 context.Context, arg1, arg2 string) (cdn.RuleSetListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByProfile", arg0, arg1, arg2)
	ret0, _ := ret[0].(cdn.RuleSetListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCDNRuleSetsClientMockRecorder) ListByProfile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByProfile", reflect.TypeOf((*MockCDNRuleSetsClient)(nil).ListByProfile), arg0, arg1, arg2)
}

type MockCDNRulesClient struct {
	ctrl     *gomock.Controller
	recorder *MockCDNRulesClientMockRecorder
}

type MockCDNRulesClientMockRecorder struct {
	mock *MockCDNRulesClient
}

func NewMockCDNRulesClient(ctrl *gomock.Controller) *MockCDNRulesClient {
	mock := &MockCDNRulesClient{ctrl: ctrl}
	mock.recorder = &MockCDNRulesClientMockRecorder{mock}
	return mock
}

func (m *MockCDNRulesClient) EXPECT() *MockCDNRulesClientMockRecorder {
	return m.recorder
}

func (m *MockCDNRulesClient) ListByRuleSet(arg0 context.Context, arg1, arg2, arg3 string) (cdn.RuleListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByRuleSet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(cdn.RuleListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCDNRulesClientMockRecorder) ListByRuleSet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByRuleSet", reflect.TypeOf((*MockCDNRulesClient)(nil).ListByRuleSet), arg0, arg1, arg2, arg3)
}

type MockCDNSecurityPoliciesClient struct {
	ctrl     *gomock.Controller
	recorder *MockCDNSecurityPoliciesClientMockRecorder
}

type MockCDNSecurityPoliciesClientMockRecorder struct {
	mock *MockCDNSecurityPoliciesClient
}

func NewMockCDNSecurityPoliciesClient(ctrl *gomock.Controller) *MockCDNSecurityPoliciesClient {
	mock := &MockCDNSecurityPoliciesClient{ctrl: ctrl}
	mock.recorder = &MockCDNSecurityPoliciesClientMockRecorder{mock}
	return mock
}

func (m *MockCDNSecurityPoliciesClient) EXPECT() *MockCDNSecurityPoliciesClientMockRecorder {
	return m.recorder
}

func (m *MockCDNSecurityPoliciesClient) ListByProfile(arg0 context.Context, arg1, arg2 string) (cdn.SecurityPolicyListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByProfile", arg0, arg1, arg2)
	ret0, _ := ret[0].(cdn.SecurityPolicyListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCDNSecurityPoliciesClientMockRecorder) ListByProfile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByProfile", reflect.TypeOf((*MockCDNSecurityPoliciesClient)(nil).ListByProfile), arg0, arg1, arg2)
}
