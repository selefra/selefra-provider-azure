package mocks

import (
	context "context"
	reflect "reflect"

	authorization "github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	gomock "github.com/golang/mock/gomock"
)

type MockAuthorizationRoleAssignmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationRoleAssignmentsClientMockRecorder
}

type MockAuthorizationRoleAssignmentsClientMockRecorder struct {
	mock *MockAuthorizationRoleAssignmentsClient
}

func NewMockAuthorizationRoleAssignmentsClient(ctrl *gomock.Controller) *MockAuthorizationRoleAssignmentsClient {
	mock := &MockAuthorizationRoleAssignmentsClient{ctrl: ctrl}
	mock.recorder = &MockAuthorizationRoleAssignmentsClientMockRecorder{mock}
	return mock
}

func (m *MockAuthorizationRoleAssignmentsClient) EXPECT() *MockAuthorizationRoleAssignmentsClientMockRecorder {
	return m.recorder
}

func (m *MockAuthorizationRoleAssignmentsClient) List(arg0 context.Context, arg1 string) (authorization.RoleAssignmentListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(authorization.RoleAssignmentListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAuthorizationRoleAssignmentsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAuthorizationRoleAssignmentsClient)(nil).List), arg0, arg1)
}

type MockAuthorizationRoleDefinitionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationRoleDefinitionsClientMockRecorder
}

type MockAuthorizationRoleDefinitionsClientMockRecorder struct {
	mock *MockAuthorizationRoleDefinitionsClient
}

func NewMockAuthorizationRoleDefinitionsClient(ctrl *gomock.Controller) *MockAuthorizationRoleDefinitionsClient {
	mock := &MockAuthorizationRoleDefinitionsClient{ctrl: ctrl}
	mock.recorder = &MockAuthorizationRoleDefinitionsClientMockRecorder{mock}
	return mock
}

func (m *MockAuthorizationRoleDefinitionsClient) EXPECT() *MockAuthorizationRoleDefinitionsClientMockRecorder {
	return m.recorder
}

func (m *MockAuthorizationRoleDefinitionsClient) List(arg0 context.Context, arg1, arg2 string) (authorization.RoleDefinitionListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(authorization.RoleDefinitionListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAuthorizationRoleDefinitionsClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAuthorizationRoleDefinitionsClient)(nil).List), arg0, arg1, arg2)
}
