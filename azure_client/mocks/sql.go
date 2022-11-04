package mocks

import (
	context "context"
	reflect "reflect"

	sql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	gomock "github.com/golang/mock/gomock"
)

type MockSQLServersClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLServersClientMockRecorder
}

type MockSQLServersClientMockRecorder struct {
	mock *MockSQLServersClient
}

func NewMockSQLServersClient(ctrl *gomock.Controller) *MockSQLServersClient {
	mock := &MockSQLServersClient{ctrl: ctrl}
	mock.recorder = &MockSQLServersClientMockRecorder{mock}
	return mock
}

func (m *MockSQLServersClient) EXPECT() *MockSQLServersClientMockRecorder {
	return m.recorder
}

func (m *MockSQLServersClient) List(arg0 context.Context) (sql.ServerListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(sql.ServerListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLServersClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSQLServersClient)(nil).List), arg0)
}

type MockSQLFirewallRulesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLFirewallRulesClientMockRecorder
}

type MockSQLFirewallRulesClientMockRecorder struct {
	mock *MockSQLFirewallRulesClient
}

func NewMockSQLFirewallRulesClient(ctrl *gomock.Controller) *MockSQLFirewallRulesClient {
	mock := &MockSQLFirewallRulesClient{ctrl: ctrl}
	mock.recorder = &MockSQLFirewallRulesClientMockRecorder{mock}
	return mock
}

func (m *MockSQLFirewallRulesClient) EXPECT() *MockSQLFirewallRulesClientMockRecorder {
	return m.recorder
}

func (m *MockSQLFirewallRulesClient) ListByServer(arg0 context.Context, arg1, arg2 string) (sql.FirewallRuleListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.FirewallRuleListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLFirewallRulesClientMockRecorder) ListByServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByServer", reflect.TypeOf((*MockSQLFirewallRulesClient)(nil).ListByServer), arg0, arg1, arg2)
}

type MockSQLServerAdminsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLServerAdminsClientMockRecorder
}

type MockSQLServerAdminsClientMockRecorder struct {
	mock *MockSQLServerAdminsClient
}

func NewMockSQLServerAdminsClient(ctrl *gomock.Controller) *MockSQLServerAdminsClient {
	mock := &MockSQLServerAdminsClient{ctrl: ctrl}
	mock.recorder = &MockSQLServerAdminsClientMockRecorder{mock}
	return mock
}

func (m *MockSQLServerAdminsClient) EXPECT() *MockSQLServerAdminsClientMockRecorder {
	return m.recorder
}

func (m *MockSQLServerAdminsClient) ListByServer(arg0 context.Context, arg1, arg2 string) (sql.AdministratorListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.AdministratorListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLServerAdminsClientMockRecorder) ListByServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByServer", reflect.TypeOf((*MockSQLServerAdminsClient)(nil).ListByServer), arg0, arg1, arg2)
}

type MockSQLServerBlobAuditingPoliciesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLServerBlobAuditingPoliciesClientMockRecorder
}

type MockSQLServerBlobAuditingPoliciesClientMockRecorder struct {
	mock *MockSQLServerBlobAuditingPoliciesClient
}

func NewMockSQLServerBlobAuditingPoliciesClient(ctrl *gomock.Controller) *MockSQLServerBlobAuditingPoliciesClient {
	mock := &MockSQLServerBlobAuditingPoliciesClient{ctrl: ctrl}
	mock.recorder = &MockSQLServerBlobAuditingPoliciesClientMockRecorder{mock}
	return mock
}

func (m *MockSQLServerBlobAuditingPoliciesClient) EXPECT() *MockSQLServerBlobAuditingPoliciesClientMockRecorder {
	return m.recorder
}

func (m *MockSQLServerBlobAuditingPoliciesClient) ListByServer(arg0 context.Context, arg1, arg2 string) (sql.ServerBlobAuditingPolicyListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.ServerBlobAuditingPolicyListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLServerBlobAuditingPoliciesClientMockRecorder) ListByServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByServer", reflect.TypeOf((*MockSQLServerBlobAuditingPoliciesClient)(nil).ListByServer), arg0, arg1, arg2)
}

type MockSQLServerDevOpsAuditingSettingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLServerDevOpsAuditingSettingsClientMockRecorder
}

type MockSQLServerDevOpsAuditingSettingsClientMockRecorder struct {
	mock *MockSQLServerDevOpsAuditingSettingsClient
}

func NewMockSQLServerDevOpsAuditingSettingsClient(ctrl *gomock.Controller) *MockSQLServerDevOpsAuditingSettingsClient {
	mock := &MockSQLServerDevOpsAuditingSettingsClient{ctrl: ctrl}
	mock.recorder = &MockSQLServerDevOpsAuditingSettingsClientMockRecorder{mock}
	return mock
}

func (m *MockSQLServerDevOpsAuditingSettingsClient) EXPECT() *MockSQLServerDevOpsAuditingSettingsClientMockRecorder {
	return m.recorder
}

func (m *MockSQLServerDevOpsAuditingSettingsClient) ListByServer(arg0 context.Context, arg1, arg2 string) (sql.ServerDevOpsAuditSettingsListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.ServerDevOpsAuditSettingsListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLServerDevOpsAuditingSettingsClientMockRecorder) ListByServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByServer", reflect.TypeOf((*MockSQLServerDevOpsAuditingSettingsClient)(nil).ListByServer), arg0, arg1, arg2)
}

type MockSQLServerVulnerabilityAssessmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLServerVulnerabilityAssessmentsClientMockRecorder
}

type MockSQLServerVulnerabilityAssessmentsClientMockRecorder struct {
	mock *MockSQLServerVulnerabilityAssessmentsClient
}

func NewMockSQLServerVulnerabilityAssessmentsClient(ctrl *gomock.Controller) *MockSQLServerVulnerabilityAssessmentsClient {
	mock := &MockSQLServerVulnerabilityAssessmentsClient{ctrl: ctrl}
	mock.recorder = &MockSQLServerVulnerabilityAssessmentsClientMockRecorder{mock}
	return mock
}

func (m *MockSQLServerVulnerabilityAssessmentsClient) EXPECT() *MockSQLServerVulnerabilityAssessmentsClientMockRecorder {
	return m.recorder
}

func (m *MockSQLServerVulnerabilityAssessmentsClient) ListByServer(arg0 context.Context, arg1, arg2 string) (sql.ServerVulnerabilityAssessmentListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.ServerVulnerabilityAssessmentListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLServerVulnerabilityAssessmentsClientMockRecorder) ListByServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByServer", reflect.TypeOf((*MockSQLServerVulnerabilityAssessmentsClient)(nil).ListByServer), arg0, arg1, arg2)
}

type MockSQLEncryptionProtectorsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLEncryptionProtectorsClientMockRecorder
}

type MockSQLEncryptionProtectorsClientMockRecorder struct {
	mock *MockSQLEncryptionProtectorsClient
}

func NewMockSQLEncryptionProtectorsClient(ctrl *gomock.Controller) *MockSQLEncryptionProtectorsClient {
	mock := &MockSQLEncryptionProtectorsClient{ctrl: ctrl}
	mock.recorder = &MockSQLEncryptionProtectorsClientMockRecorder{mock}
	return mock
}

func (m *MockSQLEncryptionProtectorsClient) EXPECT() *MockSQLEncryptionProtectorsClientMockRecorder {
	return m.recorder
}

func (m *MockSQLEncryptionProtectorsClient) Get(arg0 context.Context, arg1, arg2 string) (sql.EncryptionProtector, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.EncryptionProtector)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLEncryptionProtectorsClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSQLEncryptionProtectorsClient)(nil).Get), arg0, arg1, arg2)
}

type MockSQLVirtualNetworkRulesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLVirtualNetworkRulesClientMockRecorder
}

type MockSQLVirtualNetworkRulesClientMockRecorder struct {
	mock *MockSQLVirtualNetworkRulesClient
}

func NewMockSQLVirtualNetworkRulesClient(ctrl *gomock.Controller) *MockSQLVirtualNetworkRulesClient {
	mock := &MockSQLVirtualNetworkRulesClient{ctrl: ctrl}
	mock.recorder = &MockSQLVirtualNetworkRulesClientMockRecorder{mock}
	return mock
}

func (m *MockSQLVirtualNetworkRulesClient) EXPECT() *MockSQLVirtualNetworkRulesClientMockRecorder {
	return m.recorder
}

func (m *MockSQLVirtualNetworkRulesClient) ListByServer(arg0 context.Context, arg1, arg2 string) (sql.VirtualNetworkRuleListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.VirtualNetworkRuleListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLVirtualNetworkRulesClientMockRecorder) ListByServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByServer", reflect.TypeOf((*MockSQLVirtualNetworkRulesClient)(nil).ListByServer), arg0, arg1, arg2)
}

type MockSQLServerSecurityAlertPoliciesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLServerSecurityAlertPoliciesClientMockRecorder
}

type MockSQLServerSecurityAlertPoliciesClientMockRecorder struct {
	mock *MockSQLServerSecurityAlertPoliciesClient
}

func NewMockSQLServerSecurityAlertPoliciesClient(ctrl *gomock.Controller) *MockSQLServerSecurityAlertPoliciesClient {
	mock := &MockSQLServerSecurityAlertPoliciesClient{ctrl: ctrl}
	mock.recorder = &MockSQLServerSecurityAlertPoliciesClientMockRecorder{mock}
	return mock
}

func (m *MockSQLServerSecurityAlertPoliciesClient) EXPECT() *MockSQLServerSecurityAlertPoliciesClientMockRecorder {
	return m.recorder
}

func (m *MockSQLServerSecurityAlertPoliciesClient) ListByServer(arg0 context.Context, arg1, arg2 string) (sql.LogicalServerSecurityAlertPolicyListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.LogicalServerSecurityAlertPolicyListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLServerSecurityAlertPoliciesClientMockRecorder) ListByServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByServer", reflect.TypeOf((*MockSQLServerSecurityAlertPoliciesClient)(nil).ListByServer), arg0, arg1, arg2)
}

type MockSQLDatabasesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLDatabasesClientMockRecorder
}

type MockSQLDatabasesClientMockRecorder struct {
	mock *MockSQLDatabasesClient
}

func NewMockSQLDatabasesClient(ctrl *gomock.Controller) *MockSQLDatabasesClient {
	mock := &MockSQLDatabasesClient{ctrl: ctrl}
	mock.recorder = &MockSQLDatabasesClientMockRecorder{mock}
	return mock
}

func (m *MockSQLDatabasesClient) EXPECT() *MockSQLDatabasesClientMockRecorder {
	return m.recorder
}

func (m *MockSQLDatabasesClient) ListByServer(arg0 context.Context, arg1, arg2 string) (sql.DatabaseListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByServer", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.DatabaseListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLDatabasesClientMockRecorder) ListByServer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByServer", reflect.TypeOf((*MockSQLDatabasesClient)(nil).ListByServer), arg0, arg1, arg2)
}

type MockSQLDatabaseBlobAuditingPoliciesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLDatabaseBlobAuditingPoliciesClientMockRecorder
}

type MockSQLDatabaseBlobAuditingPoliciesClientMockRecorder struct {
	mock *MockSQLDatabaseBlobAuditingPoliciesClient
}

func NewMockSQLDatabaseBlobAuditingPoliciesClient(ctrl *gomock.Controller) *MockSQLDatabaseBlobAuditingPoliciesClient {
	mock := &MockSQLDatabaseBlobAuditingPoliciesClient{ctrl: ctrl}
	mock.recorder = &MockSQLDatabaseBlobAuditingPoliciesClientMockRecorder{mock}
	return mock
}

func (m *MockSQLDatabaseBlobAuditingPoliciesClient) EXPECT() *MockSQLDatabaseBlobAuditingPoliciesClientMockRecorder {
	return m.recorder
}

func (m *MockSQLDatabaseBlobAuditingPoliciesClient) ListByDatabase(arg0 context.Context, arg1, arg2, arg3 string) (sql.DatabaseBlobAuditingPolicyListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByDatabase", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.DatabaseBlobAuditingPolicyListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLDatabaseBlobAuditingPoliciesClientMockRecorder) ListByDatabase(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByDatabase", reflect.TypeOf((*MockSQLDatabaseBlobAuditingPoliciesClient)(nil).ListByDatabase), arg0, arg1, arg2, arg3)
}

type MockSQLDatabaseThreatDetectionPoliciesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLDatabaseThreatDetectionPoliciesClientMockRecorder
}

type MockSQLDatabaseThreatDetectionPoliciesClientMockRecorder struct {
	mock *MockSQLDatabaseThreatDetectionPoliciesClient
}

func NewMockSQLDatabaseThreatDetectionPoliciesClient(ctrl *gomock.Controller) *MockSQLDatabaseThreatDetectionPoliciesClient {
	mock := &MockSQLDatabaseThreatDetectionPoliciesClient{ctrl: ctrl}
	mock.recorder = &MockSQLDatabaseThreatDetectionPoliciesClientMockRecorder{mock}
	return mock
}

func (m *MockSQLDatabaseThreatDetectionPoliciesClient) EXPECT() *MockSQLDatabaseThreatDetectionPoliciesClientMockRecorder {
	return m.recorder
}

func (m *MockSQLDatabaseThreatDetectionPoliciesClient) Get(arg0 context.Context, arg1, arg2, arg3 string) (sql.DatabaseSecurityAlertPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.DatabaseSecurityAlertPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLDatabaseThreatDetectionPoliciesClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSQLDatabaseThreatDetectionPoliciesClient)(nil).Get), arg0, arg1, arg2, arg3)
}

type MockSQLDatabaseVulnerabilityAssessmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLDatabaseVulnerabilityAssessmentsClientMockRecorder
}

type MockSQLDatabaseVulnerabilityAssessmentsClientMockRecorder struct {
	mock *MockSQLDatabaseVulnerabilityAssessmentsClient
}

func NewMockSQLDatabaseVulnerabilityAssessmentsClient(ctrl *gomock.Controller) *MockSQLDatabaseVulnerabilityAssessmentsClient {
	mock := &MockSQLDatabaseVulnerabilityAssessmentsClient{ctrl: ctrl}
	mock.recorder = &MockSQLDatabaseVulnerabilityAssessmentsClientMockRecorder{mock}
	return mock
}

func (m *MockSQLDatabaseVulnerabilityAssessmentsClient) EXPECT() *MockSQLDatabaseVulnerabilityAssessmentsClientMockRecorder {
	return m.recorder
}

func (m *MockSQLDatabaseVulnerabilityAssessmentsClient) ListByDatabase(arg0 context.Context, arg1, arg2, arg3 string) (sql.DatabaseVulnerabilityAssessmentListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByDatabase", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.DatabaseVulnerabilityAssessmentListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLDatabaseVulnerabilityAssessmentsClientMockRecorder) ListByDatabase(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByDatabase", reflect.TypeOf((*MockSQLDatabaseVulnerabilityAssessmentsClient)(nil).ListByDatabase), arg0, arg1, arg2, arg3)
}

type MockSQLDatabaseVulnerabilityAssessmentScansClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLDatabaseVulnerabilityAssessmentScansClientMockRecorder
}

type MockSQLDatabaseVulnerabilityAssessmentScansClientMockRecorder struct {
	mock *MockSQLDatabaseVulnerabilityAssessmentScansClient
}

func NewMockSQLDatabaseVulnerabilityAssessmentScansClient(ctrl *gomock.Controller) *MockSQLDatabaseVulnerabilityAssessmentScansClient {
	mock := &MockSQLDatabaseVulnerabilityAssessmentScansClient{ctrl: ctrl}
	mock.recorder = &MockSQLDatabaseVulnerabilityAssessmentScansClientMockRecorder{mock}
	return mock
}

func (m *MockSQLDatabaseVulnerabilityAssessmentScansClient) EXPECT() *MockSQLDatabaseVulnerabilityAssessmentScansClientMockRecorder {
	return m.recorder
}

func (m *MockSQLDatabaseVulnerabilityAssessmentScansClient) ListByDatabase(arg0 context.Context, arg1, arg2, arg3 string) (sql.VulnerabilityAssessmentScanRecordListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByDatabase", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.VulnerabilityAssessmentScanRecordListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLDatabaseVulnerabilityAssessmentScansClientMockRecorder) ListByDatabase(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByDatabase", reflect.TypeOf((*MockSQLDatabaseVulnerabilityAssessmentScansClient)(nil).ListByDatabase), arg0, arg1, arg2, arg3)
}

type MockSQLTransparentDataEncryptionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLTransparentDataEncryptionsClientMockRecorder
}

type MockSQLTransparentDataEncryptionsClientMockRecorder struct {
	mock *MockSQLTransparentDataEncryptionsClient
}

func NewMockSQLTransparentDataEncryptionsClient(ctrl *gomock.Controller) *MockSQLTransparentDataEncryptionsClient {
	mock := &MockSQLTransparentDataEncryptionsClient{ctrl: ctrl}
	mock.recorder = &MockSQLTransparentDataEncryptionsClientMockRecorder{mock}
	return mock
}

func (m *MockSQLTransparentDataEncryptionsClient) EXPECT() *MockSQLTransparentDataEncryptionsClientMockRecorder {
	return m.recorder
}

func (m *MockSQLTransparentDataEncryptionsClient) Get(arg0 context.Context, arg1, arg2, arg3 string) (sql.TransparentDataEncryption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.TransparentDataEncryption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLTransparentDataEncryptionsClientMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSQLTransparentDataEncryptionsClient)(nil).Get), arg0, arg1, arg2, arg3)
}

type MockSQLBackupLongTermRetentionPoliciesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLBackupLongTermRetentionPoliciesClientMockRecorder
}

type MockSQLBackupLongTermRetentionPoliciesClientMockRecorder struct {
	mock *MockSQLBackupLongTermRetentionPoliciesClient
}

func NewMockSQLBackupLongTermRetentionPoliciesClient(ctrl *gomock.Controller) *MockSQLBackupLongTermRetentionPoliciesClient {
	mock := &MockSQLBackupLongTermRetentionPoliciesClient{ctrl: ctrl}
	mock.recorder = &MockSQLBackupLongTermRetentionPoliciesClientMockRecorder{mock}
	return mock
}

func (m *MockSQLBackupLongTermRetentionPoliciesClient) EXPECT() *MockSQLBackupLongTermRetentionPoliciesClientMockRecorder {
	return m.recorder
}

func (m *MockSQLBackupLongTermRetentionPoliciesClient) ListByDatabase(arg0 context.Context, arg1, arg2, arg3 string) (sql.BackupLongTermRetentionPolicy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByDatabase", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.BackupLongTermRetentionPolicy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLBackupLongTermRetentionPoliciesClientMockRecorder) ListByDatabase(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByDatabase", reflect.TypeOf((*MockSQLBackupLongTermRetentionPoliciesClient)(nil).ListByDatabase), arg0, arg1, arg2, arg3)
}

type MockSQLManagedInstancesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLManagedInstancesClientMockRecorder
}

type MockSQLManagedInstancesClientMockRecorder struct {
	mock *MockSQLManagedInstancesClient
}

func NewMockSQLManagedInstancesClient(ctrl *gomock.Controller) *MockSQLManagedInstancesClient {
	mock := &MockSQLManagedInstancesClient{ctrl: ctrl}
	mock.recorder = &MockSQLManagedInstancesClientMockRecorder{mock}
	return mock
}

func (m *MockSQLManagedInstancesClient) EXPECT() *MockSQLManagedInstancesClientMockRecorder {
	return m.recorder
}

func (m *MockSQLManagedInstancesClient) List(arg0 context.Context) (sql.ManagedInstanceListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(sql.ManagedInstanceListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLManagedInstancesClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSQLManagedInstancesClient)(nil).List), arg0)
}

type MockSQLManagedInstanceVulnerabilityAssessmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLManagedInstanceVulnerabilityAssessmentsClientMockRecorder
}

type MockSQLManagedInstanceVulnerabilityAssessmentsClientMockRecorder struct {
	mock *MockSQLManagedInstanceVulnerabilityAssessmentsClient
}

func NewMockSQLManagedInstanceVulnerabilityAssessmentsClient(ctrl *gomock.Controller) *MockSQLManagedInstanceVulnerabilityAssessmentsClient {
	mock := &MockSQLManagedInstanceVulnerabilityAssessmentsClient{ctrl: ctrl}
	mock.recorder = &MockSQLManagedInstanceVulnerabilityAssessmentsClientMockRecorder{mock}
	return mock
}

func (m *MockSQLManagedInstanceVulnerabilityAssessmentsClient) EXPECT() *MockSQLManagedInstanceVulnerabilityAssessmentsClientMockRecorder {
	return m.recorder
}

func (m *MockSQLManagedInstanceVulnerabilityAssessmentsClient) ListByInstance(arg0 context.Context, arg1, arg2 string) (sql.ManagedInstanceVulnerabilityAssessmentListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByInstance", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.ManagedInstanceVulnerabilityAssessmentListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLManagedInstanceVulnerabilityAssessmentsClientMockRecorder) ListByInstance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByInstance", reflect.TypeOf((*MockSQLManagedInstanceVulnerabilityAssessmentsClient)(nil).ListByInstance), arg0, arg1, arg2)
}

type MockSQLManagedInstanceEncryptionProtectorsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLManagedInstanceEncryptionProtectorsClientMockRecorder
}

type MockSQLManagedInstanceEncryptionProtectorsClientMockRecorder struct {
	mock *MockSQLManagedInstanceEncryptionProtectorsClient
}

func NewMockSQLManagedInstanceEncryptionProtectorsClient(ctrl *gomock.Controller) *MockSQLManagedInstanceEncryptionProtectorsClient {
	mock := &MockSQLManagedInstanceEncryptionProtectorsClient{ctrl: ctrl}
	mock.recorder = &MockSQLManagedInstanceEncryptionProtectorsClientMockRecorder{mock}
	return mock
}

func (m *MockSQLManagedInstanceEncryptionProtectorsClient) EXPECT() *MockSQLManagedInstanceEncryptionProtectorsClientMockRecorder {
	return m.recorder
}

func (m *MockSQLManagedInstanceEncryptionProtectorsClient) ListByInstance(arg0 context.Context, arg1, arg2 string) (sql.ManagedInstanceEncryptionProtectorListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByInstance", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.ManagedInstanceEncryptionProtectorListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLManagedInstanceEncryptionProtectorsClientMockRecorder) ListByInstance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByInstance", reflect.TypeOf((*MockSQLManagedInstanceEncryptionProtectorsClient)(nil).ListByInstance), arg0, arg1, arg2)
}

type MockSQLManagedDatabasesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLManagedDatabasesClientMockRecorder
}

type MockSQLManagedDatabasesClientMockRecorder struct {
	mock *MockSQLManagedDatabasesClient
}

func NewMockSQLManagedDatabasesClient(ctrl *gomock.Controller) *MockSQLManagedDatabasesClient {
	mock := &MockSQLManagedDatabasesClient{ctrl: ctrl}
	mock.recorder = &MockSQLManagedDatabasesClientMockRecorder{mock}
	return mock
}

func (m *MockSQLManagedDatabasesClient) EXPECT() *MockSQLManagedDatabasesClientMockRecorder {
	return m.recorder
}

func (m *MockSQLManagedDatabasesClient) ListByInstance(arg0 context.Context, arg1, arg2 string) (sql.ManagedDatabaseListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByInstance", arg0, arg1, arg2)
	ret0, _ := ret[0].(sql.ManagedDatabaseListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLManagedDatabasesClientMockRecorder) ListByInstance(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByInstance", reflect.TypeOf((*MockSQLManagedDatabasesClient)(nil).ListByInstance), arg0, arg1, arg2)
}

type MockSQLManagedDatabaseVulnerabilityAssessmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLManagedDatabaseVulnerabilityAssessmentsClientMockRecorder
}

type MockSQLManagedDatabaseVulnerabilityAssessmentsClientMockRecorder struct {
	mock *MockSQLManagedDatabaseVulnerabilityAssessmentsClient
}

func NewMockSQLManagedDatabaseVulnerabilityAssessmentsClient(ctrl *gomock.Controller) *MockSQLManagedDatabaseVulnerabilityAssessmentsClient {
	mock := &MockSQLManagedDatabaseVulnerabilityAssessmentsClient{ctrl: ctrl}
	mock.recorder = &MockSQLManagedDatabaseVulnerabilityAssessmentsClientMockRecorder{mock}
	return mock
}

func (m *MockSQLManagedDatabaseVulnerabilityAssessmentsClient) EXPECT() *MockSQLManagedDatabaseVulnerabilityAssessmentsClientMockRecorder {
	return m.recorder
}

func (m *MockSQLManagedDatabaseVulnerabilityAssessmentsClient) ListByDatabase(arg0 context.Context, arg1, arg2, arg3 string) (sql.DatabaseVulnerabilityAssessmentListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByDatabase", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.DatabaseVulnerabilityAssessmentListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLManagedDatabaseVulnerabilityAssessmentsClientMockRecorder) ListByDatabase(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByDatabase", reflect.TypeOf((*MockSQLManagedDatabaseVulnerabilityAssessmentsClient)(nil).ListByDatabase), arg0, arg1, arg2, arg3)
}

type MockSQLManagedDatabaseVulnerabilityAssessmentScansClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLManagedDatabaseVulnerabilityAssessmentScansClientMockRecorder
}

type MockSQLManagedDatabaseVulnerabilityAssessmentScansClientMockRecorder struct {
	mock *MockSQLManagedDatabaseVulnerabilityAssessmentScansClient
}

func NewMockSQLManagedDatabaseVulnerabilityAssessmentScansClient(ctrl *gomock.Controller) *MockSQLManagedDatabaseVulnerabilityAssessmentScansClient {
	mock := &MockSQLManagedDatabaseVulnerabilityAssessmentScansClient{ctrl: ctrl}
	mock.recorder = &MockSQLManagedDatabaseVulnerabilityAssessmentScansClientMockRecorder{mock}
	return mock
}

func (m *MockSQLManagedDatabaseVulnerabilityAssessmentScansClient) EXPECT() *MockSQLManagedDatabaseVulnerabilityAssessmentScansClientMockRecorder {
	return m.recorder
}

func (m *MockSQLManagedDatabaseVulnerabilityAssessmentScansClient) ListByDatabase(arg0 context.Context, arg1, arg2, arg3 string) (sql.VulnerabilityAssessmentScanRecordListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByDatabase", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(sql.VulnerabilityAssessmentScanRecordListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSQLManagedDatabaseVulnerabilityAssessmentScansClientMockRecorder) ListByDatabase(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByDatabase", reflect.TypeOf((*MockSQLManagedDatabaseVulnerabilityAssessmentScansClient)(nil).ListByDatabase), arg0, arg1, arg2, arg3)
}
