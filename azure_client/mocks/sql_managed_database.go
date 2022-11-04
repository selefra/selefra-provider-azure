package mocks

import (
	context "context"
	reflect "reflect"

	sql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	gomock "github.com/golang/mock/gomock"
)

type MockManagedDatabasesClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagedDatabasesClientMockRecorder
}

type MockManagedDatabasesClientMockRecorder struct {
	mock *MockManagedDatabasesClient
}

func NewMockManagedDatabasesClient(ctrl *gomock.Controller) *MockManagedDatabasesClient {
	mock := &MockManagedDatabasesClient{ctrl: ctrl}
	mock.recorder = &MockManagedDatabasesClientMockRecorder{mock}
	return mock
}

func (m *MockManagedDatabasesClient) EXPECT() *MockManagedDatabasesClientMockRecorder {
	return m.recorder
}

func (m *MockManagedDatabasesClient) ListByInstance(arg0 context.Context, arg1, arg2 string) (sql.ManagedDatabaseListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByInstance", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.ManagedDatabaseListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockManagedDatabasesClientMockRecorder) ListByInstance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByInstance", reflect.TypeOf((*MockManagedDatabasesClient)(nil).ListByInstance), arg0, arg1, arg2)
}

type MockManagedDatabaseVulnerabilityAssessmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagedDatabaseVulnerabilityAssessmentsClientMockRecorder
}

type MockManagedDatabaseVulnerabilityAssessmentsClientMockRecorder struct {
	mock *MockManagedDatabaseVulnerabilityAssessmentsClient
}

func NewMockManagedDatabaseVulnerabilityAssessmentsClient(ctrl *gomock.Controller) *MockManagedDatabaseVulnerabilityAssessmentsClient {
	mock := &MockManagedDatabaseVulnerabilityAssessmentsClient{ctrl: ctrl}
	mock.recorder = &MockManagedDatabaseVulnerabilityAssessmentsClientMockRecorder{mock}
	return mock
}

func (m *MockManagedDatabaseVulnerabilityAssessmentsClient) EXPECT() *MockManagedDatabaseVulnerabilityAssessmentsClientMockRecorder {
	return m.recorder
}

func (m *MockManagedDatabaseVulnerabilityAssessmentsClient) ListByDatabase(arg0 context.Context, arg1, arg2, arg3 string) (sql.DatabaseVulnerabilityAssessmentListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByDatabase", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.DatabaseVulnerabilityAssessmentListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockManagedDatabaseVulnerabilityAssessmentsClientMockRecorder) ListByDatabase(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByDatabase", reflect.TypeOf((*MockManagedDatabaseVulnerabilityAssessmentsClient)(nil).ListByDatabase), arg0, arg1, arg2, arg3)
}

type MockManagedDatabaseVulnerabilityAssessmentScansClient struct {
	ctrl     *gomock.Controller
	recorder *MockManagedDatabaseVulnerabilityAssessmentScansClientMockRecorder
}

type MockManagedDatabaseVulnerabilityAssessmentScansClientMockRecorder struct {
	mock *MockManagedDatabaseVulnerabilityAssessmentScansClient
}

func NewMockManagedDatabaseVulnerabilityAssessmentScansClient(ctrl *gomock.Controller) *MockManagedDatabaseVulnerabilityAssessmentScansClient {
	mock := &MockManagedDatabaseVulnerabilityAssessmentScansClient{ctrl: ctrl}
	mock.recorder = &MockManagedDatabaseVulnerabilityAssessmentScansClientMockRecorder{mock}
	return mock
}

func (m *MockManagedDatabaseVulnerabilityAssessmentScansClient) EXPECT() *MockManagedDatabaseVulnerabilityAssessmentScansClientMockRecorder {
	return m.recorder
}

func (m *MockManagedDatabaseVulnerabilityAssessmentScansClient) ListByDatabase(arg0 context.Context, arg1, arg2, arg3 string) (sql.VulnerabilityAssessmentScanRecordListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByDatabase", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.VulnerabilityAssessmentScanRecordListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockManagedDatabaseVulnerabilityAssessmentScansClientMockRecorder) ListByDatabase(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByDatabase", reflect.TypeOf((*MockManagedDatabaseVulnerabilityAssessmentScansClient)(nil).ListByDatabase), arg0, arg1, arg2, arg3)
}
