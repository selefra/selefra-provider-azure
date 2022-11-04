package mocks

import (
	context "context"
	reflect "reflect"

	insights "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	insights0 "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"
	resources "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	gomock "github.com/golang/mock/gomock"
)

type MockMonitorActivityLogAlertsClient struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorActivityLogAlertsClientMockRecorder
}

type MockMonitorActivityLogAlertsClientMockRecorder struct {
	mock *MockMonitorActivityLogAlertsClient
}

func NewMockMonitorActivityLogAlertsClient(ctrl *gomock.Controller) *MockMonitorActivityLogAlertsClient {
	mock := &MockMonitorActivityLogAlertsClient{ctrl: ctrl}
	mock.recorder = &MockMonitorActivityLogAlertsClientMockRecorder{mock}
	return mock
}

func (m *MockMonitorActivityLogAlertsClient) EXPECT() *MockMonitorActivityLogAlertsClientMockRecorder {
	return m.recorder
}

func (m *MockMonitorActivityLogAlertsClient) ListBySubscriptionID(arg0 context.Context) (insights.ActivityLogAlertList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscriptionID", arg0)
	ret0, _ := ret[0].(insights.ActivityLogAlertList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMonitorActivityLogAlertsClientMockRecorder) ListBySubscriptionID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscriptionID", reflect.TypeOf((*MockMonitorActivityLogAlertsClient)(nil).ListBySubscriptionID), arg0)
}

type MockMonitorLogProfilesClient struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorLogProfilesClientMockRecorder
}

type MockMonitorLogProfilesClientMockRecorder struct {
	mock *MockMonitorLogProfilesClient
}

func NewMockMonitorLogProfilesClient(ctrl *gomock.Controller) *MockMonitorLogProfilesClient {
	mock := &MockMonitorLogProfilesClient{ctrl: ctrl}
	mock.recorder = &MockMonitorLogProfilesClientMockRecorder{mock}
	return mock
}

func (m *MockMonitorLogProfilesClient) EXPECT() *MockMonitorLogProfilesClientMockRecorder {
	return m.recorder
}

func (m *MockMonitorLogProfilesClient) List(arg0 context.Context) (insights0.LogProfileCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(insights0.LogProfileCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMonitorLogProfilesClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMonitorLogProfilesClient)(nil).List), arg0)
}

type MockMonitorDiagnosticSettingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorDiagnosticSettingsClientMockRecorder
}

type MockMonitorDiagnosticSettingsClientMockRecorder struct {
	mock *MockMonitorDiagnosticSettingsClient
}

func NewMockMonitorDiagnosticSettingsClient(ctrl *gomock.Controller) *MockMonitorDiagnosticSettingsClient {
	mock := &MockMonitorDiagnosticSettingsClient{ctrl: ctrl}
	mock.recorder = &MockMonitorDiagnosticSettingsClientMockRecorder{mock}
	return mock
}

func (m *MockMonitorDiagnosticSettingsClient) EXPECT() *MockMonitorDiagnosticSettingsClientMockRecorder {
	return m.recorder
}

func (m *MockMonitorDiagnosticSettingsClient) List(arg0 context.Context, arg1 string) (insights0.DiagnosticSettingsResourceCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(insights0.DiagnosticSettingsResourceCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMonitorDiagnosticSettingsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMonitorDiagnosticSettingsClient)(nil).List), arg0, arg1)
}

type MockMonitorActivityLogsClient struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorActivityLogsClientMockRecorder
}

type MockMonitorActivityLogsClientMockRecorder struct {
	mock *MockMonitorActivityLogsClient
}

func NewMockMonitorActivityLogsClient(ctrl *gomock.Controller) *MockMonitorActivityLogsClient {
	mock := &MockMonitorActivityLogsClient{ctrl: ctrl}
	mock.recorder = &MockMonitorActivityLogsClientMockRecorder{mock}
	return mock
}

func (m *MockMonitorActivityLogsClient) EXPECT() *MockMonitorActivityLogsClientMockRecorder {
	return m.recorder
}

func (m *MockMonitorActivityLogsClient) List(arg0 context.Context, arg1, arg2 string) (insights0.EventDataCollectionPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(insights0.EventDataCollectionPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMonitorActivityLogsClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMonitorActivityLogsClient)(nil).List), arg0, arg1, arg2)
}

type MockMonitorResourcesClient struct {
	ctrl     *gomock.Controller
	recorder *MockMonitorResourcesClientMockRecorder
}

type MockMonitorResourcesClientMockRecorder struct {
	mock *MockMonitorResourcesClient
}

func NewMockMonitorResourcesClient(ctrl *gomock.Controller) *MockMonitorResourcesClient {
	mock := &MockMonitorResourcesClient{ctrl: ctrl}
	mock.recorder = &MockMonitorResourcesClientMockRecorder{mock}
	return mock
}

func (m *MockMonitorResourcesClient) EXPECT() *MockMonitorResourcesClientMockRecorder {
	return m.recorder
}

func (m *MockMonitorResourcesClient) List(arg0 context.Context, arg1, arg2 string, arg3 *int32) (resources.ListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(resources.ListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMonitorResourcesClientMockRecorder) List(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMonitorResourcesClient)(nil).List), arg0, arg1, arg2, arg3)
}
