package mocks

import (
	context "context"
	reflect "reflect"

	policy "github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2020-03-01-preview/policy"
	links "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/links"
	resources "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	gomock "github.com/golang/mock/gomock"
)

type MockResourcesGroupsClient struct {
	ctrl     *gomock.Controller
	recorder *MockResourcesGroupsClientMockRecorder
}

type MockResourcesGroupsClientMockRecorder struct {
	mock *MockResourcesGroupsClient
}

func NewMockResourcesGroupsClient(ctrl *gomock.Controller) *MockResourcesGroupsClient {
	mock := &MockResourcesGroupsClient{ctrl: ctrl}
	mock.recorder = &MockResourcesGroupsClientMockRecorder{mock}
	return mock
}

func (m *MockResourcesGroupsClient) EXPECT() *MockResourcesGroupsClientMockRecorder {
	return m.recorder
}

func (m *MockResourcesGroupsClient) List(arg0 context.Context, arg1 string, arg2 *int32) (resources.GroupListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(resources.GroupListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourcesGroupsClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockResourcesGroupsClient)(nil).List), arg0, arg1, arg2)
}

type MockResourcesPolicyAssignmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockResourcesPolicyAssignmentsClientMockRecorder
}

type MockResourcesPolicyAssignmentsClientMockRecorder struct {
	mock *MockResourcesPolicyAssignmentsClient
}

func NewMockResourcesPolicyAssignmentsClient(ctrl *gomock.Controller) *MockResourcesPolicyAssignmentsClient {
	mock := &MockResourcesPolicyAssignmentsClient{ctrl: ctrl}
	mock.recorder = &MockResourcesPolicyAssignmentsClientMockRecorder{mock}
	return mock
}

func (m *MockResourcesPolicyAssignmentsClient) EXPECT() *MockResourcesPolicyAssignmentsClientMockRecorder {
	return m.recorder
}

func (m *MockResourcesPolicyAssignmentsClient) List(arg0 context.Context, arg1, arg2 string, arg3 *int32) (policy.AssignmentListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(policy.AssignmentListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourcesPolicyAssignmentsClientMockRecorder) List(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockResourcesPolicyAssignmentsClient)(nil).List), arg0, arg1, arg2, arg3)
}

type MockResourcesLinksClient struct {
	ctrl     *gomock.Controller
	recorder *MockResourcesLinksClientMockRecorder
}

type MockResourcesLinksClientMockRecorder struct {
	mock *MockResourcesLinksClient
}

func NewMockResourcesLinksClient(ctrl *gomock.Controller) *MockResourcesLinksClient {
	mock := &MockResourcesLinksClient{ctrl: ctrl}
	mock.recorder = &MockResourcesLinksClientMockRecorder{mock}
	return mock
}

func (m *MockResourcesLinksClient) EXPECT() *MockResourcesLinksClientMockRecorder {
	return m.recorder
}

func (m *MockResourcesLinksClient) ListAtSubscription(arg0 context.Context, arg1 string) (links.ResourceLinkResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAtSubscription", arg0, arg1)
	ret0, _ := ret[0].(links.ResourceLinkResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockResourcesLinksClientMockRecorder) ListAtSubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAtSubscription", reflect.TypeOf((*MockResourcesLinksClient)(nil).ListAtSubscription), arg0, arg1)
}
