package mocks

import (
	context "context"
	reflect "reflect"

	sql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	gomock "github.com/golang/mock/gomock"
)

type MockManagedInstancesClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagedInstancesClientMockRecorder
}

type MockManagedInstancesClientMockRecorder struct {
	mock *MockManagedInstancesClient
}

func NewMockManagedInstancesClient(ctrl *gomock.Controller) *MockManagedInstancesClient {
	mock := &MockManagedInstancesClient{ctrl: ctrl}
	mock.recorder = &MockManagedInstancesClientMockRecorder{mock}
	return mock
}

func (m *MockManagedInstancesClient) EXPECT() *MockManagedInstancesClientMockRecorder {
	return m.recorder
}

func (m *MockManagedInstancesClient) List(arg0 context.Context) (sql.ManagedInstanceListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(sql.ManagedInstanceListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockManagedInstancesClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockManagedInstancesClient)(nil).List), arg0)
}

type MockManagedInstanceVulnerabilityAssessmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagedInstanceVulnerabilityAssessmentsClientMockRecorder
}

type MockManagedInstanceVulnerabilityAssessmentsClientMockRecorder struct {
	mock *MockManagedInstanceVulnerabilityAssessmentsClient
}

func NewMockManagedInstanceVulnerabilityAssessmentsClient(ctrl *gomock.Controller) *MockManagedInstanceVulnerabilityAssessmentsClient {
	mock := &MockManagedInstanceVulnerabilityAssessmentsClient{ctrl: ctrl}
	mock.recorder = &MockManagedInstanceVulnerabilityAssessmentsClientMockRecorder{mock}
	return mock
}

func (m *MockManagedInstanceVulnerabilityAssessmentsClient) EXPECT() *MockManagedInstanceVulnerabilityAssessmentsClientMockRecorder {
	return m.recorder
}

func (m *MockManagedInstanceVulnerabilityAssessmentsClient) ListByInstance(arg0 context.Context, arg1, arg2 string) (sql.ManagedInstanceVulnerabilityAssessmentListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByInstance", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.ManagedInstanceVulnerabilityAssessmentListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockManagedInstanceVulnerabilityAssessmentsClientMockRecorder) ListByInstance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByInstance", reflect.TypeOf((*MockManagedInstanceVulnerabilityAssessmentsClient)(nil).ListByInstance), arg0, arg1, arg2)
}

type MockManagedInstanceEncryptionProtectorsClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagedInstanceEncryptionProtectorsClientMockRecorder
}

type MockManagedInstanceEncryptionProtectorsClientMockRecorder struct {
	mock *MockManagedInstanceEncryptionProtectorsClient
}

func NewMockManagedInstanceEncryptionProtectorsClient(ctrl *gomock.Controller) *MockManagedInstanceEncryptionProtectorsClient {
	mock := &MockManagedInstanceEncryptionProtectorsClient{ctrl: ctrl}
	mock.recorder = &MockManagedInstanceEncryptionProtectorsClientMockRecorder{mock}
	return mock
}

func (m *MockManagedInstanceEncryptionProtectorsClient) EXPECT() *MockManagedInstanceEncryptionProtectorsClientMockRecorder {
	return m.recorder
}

func (m *MockManagedInstanceEncryptionProtectorsClient) ListByInstance(arg0 context.Context, arg1, arg2 string) (sql.ManagedInstanceEncryptionProtectorListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByInstance", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.ManagedInstanceEncryptionProtectorListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockManagedInstanceEncryptionProtectorsClientMockRecorder) ListByInstance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByInstance", reflect.TypeOf((*MockManagedInstanceEncryptionProtectorsClient)(nil).ListByInstance), arg0, arg1, arg2)
}
