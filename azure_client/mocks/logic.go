package mocks

import (
	context "context"
	reflect "reflect"

	logic "github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2019-05-01/logic"
	insights "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-06-01/insights"
	gomock "github.com/golang/mock/gomock"
)

type MockLogicDiagnosticSettingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockLogicDiagnosticSettingsClientMockRecorder
}

type MockLogicDiagnosticSettingsClientMockRecorder struct {
	mock *MockLogicDiagnosticSettingsClient
}

func NewMockLogicDiagnosticSettingsClient(ctrl *gomock.Controller) *MockLogicDiagnosticSettingsClient {
	mock := &MockLogicDiagnosticSettingsClient{ctrl: ctrl}
	mock.recorder = &MockLogicDiagnosticSettingsClientMockRecorder{mock}
	return mock
}

func (m *MockLogicDiagnosticSettingsClient) EXPECT() *MockLogicDiagnosticSettingsClientMockRecorder {
	return m.recorder
}

func (m *MockLogicDiagnosticSettingsClient) List(arg0 context.Context, arg1 string) (insights.DiagnosticSettingsResourceCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(insights.DiagnosticSettingsResourceCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLogicDiagnosticSettingsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLogicDiagnosticSettingsClient)(nil).List), arg0, arg1)
}

type MockLogicWorkflowsClient struct {
	ctrl     *gomock.Controller
	recorder *MockLogicWorkflowsClientMockRecorder
}

type MockLogicWorkflowsClientMockRecorder struct {
	mock *MockLogicWorkflowsClient
}

func NewMockLogicWorkflowsClient(ctrl *gomock.Controller) *MockLogicWorkflowsClient {
	mock := &MockLogicWorkflowsClient{ctrl: ctrl}
	mock.recorder = &MockLogicWorkflowsClientMockRecorder{mock}
	return mock
}

func (m *MockLogicWorkflowsClient) EXPECT() *MockLogicWorkflowsClientMockRecorder {
	return m.recorder
}

func (m *MockLogicWorkflowsClient) ListBySubscription(arg0 context.Context, arg1 *int32, arg2 string) (logic.WorkflowListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscription", arg0, arg1, arg2)
	ret0, _ := ret[0].(logic.WorkflowListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockLogicWorkflowsClientMockRecorder) ListBySubscription(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscription", reflect.TypeOf((*MockLogicWorkflowsClient)(nil).ListBySubscription), arg0, arg1, arg2)
}
