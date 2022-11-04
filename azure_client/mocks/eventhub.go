package mocks

import (
	context "context"
	reflect "reflect"

	eventhub "github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"
	gomock "github.com/golang/mock/gomock"
)

type MockEventHubNamespacesClient struct {
	ctrl     *gomock.Controller
	recorder *MockEventHubNamespacesClientMockRecorder
}

type MockEventHubNamespacesClientMockRecorder struct {
	mock *MockEventHubNamespacesClient
}

func NewMockEventHubNamespacesClient(ctrl *gomock.Controller) *MockEventHubNamespacesClient {
	mock := &MockEventHubNamespacesClient{ctrl: ctrl}
	mock.recorder = &MockEventHubNamespacesClientMockRecorder{mock}
	return mock
}

func (m *MockEventHubNamespacesClient) EXPECT() *MockEventHubNamespacesClientMockRecorder {
	return m.recorder
}

func (m *MockEventHubNamespacesClient) List(arg0 context.Context) (eventhub.EHNamespaceListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(eventhub.EHNamespaceListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventHubNamespacesClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockEventHubNamespacesClient)(nil).List), arg0)
}

type MockEventHubNetworkRuleSetsClient struct {
	ctrl     *gomock.Controller
	recorder *MockEventHubNetworkRuleSetsClientMockRecorder
}

type MockEventHubNetworkRuleSetsClientMockRecorder struct {
	mock *MockEventHubNetworkRuleSetsClient
}

func NewMockEventHubNetworkRuleSetsClient(ctrl *gomock.Controller) *MockEventHubNetworkRuleSetsClient {
	mock := &MockEventHubNetworkRuleSetsClient{ctrl: ctrl}
	mock.recorder = &MockEventHubNetworkRuleSetsClientMockRecorder{mock}
	return mock
}

func (m *MockEventHubNetworkRuleSetsClient) EXPECT() *MockEventHubNetworkRuleSetsClientMockRecorder {
	return m.recorder
}

func (m *MockEventHubNetworkRuleSetsClient) GetNetworkRuleSet(arg0 context.Context, arg1, arg2 string) (eventhub.NetworkRuleSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkRuleSet", arg0, arg1, arg2)
	ret0, _ := ret[0].(eventhub.NetworkRuleSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockEventHubNetworkRuleSetsClientMockRecorder) GetNetworkRuleSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkRuleSet", reflect.TypeOf((*MockEventHubNetworkRuleSetsClient)(nil).GetNetworkRuleSet), arg0, arg1, arg2)
}
