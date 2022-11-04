package mocks

import (
	context "context"
	reflect "reflect"

	security "github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	gomock "github.com/golang/mock/gomock"
)

type MockSecurityAutoProvisioningSettingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecurityAutoProvisioningSettingsClientMockRecorder
}

type MockSecurityAutoProvisioningSettingsClientMockRecorder struct {
	mock *MockSecurityAutoProvisioningSettingsClient
}

func NewMockSecurityAutoProvisioningSettingsClient(ctrl *gomock.Controller) *MockSecurityAutoProvisioningSettingsClient {
	mock := &MockSecurityAutoProvisioningSettingsClient{ctrl: ctrl}
	mock.recorder = &MockSecurityAutoProvisioningSettingsClientMockRecorder{mock}
	return mock
}

func (m *MockSecurityAutoProvisioningSettingsClient) EXPECT() *MockSecurityAutoProvisioningSettingsClientMockRecorder {
	return m.recorder
}

func (m *MockSecurityAutoProvisioningSettingsClient) List(arg0 context.Context) (security.AutoProvisioningSettingListPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(security.AutoProvisioningSettingListPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityAutoProvisioningSettingsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecurityAutoProvisioningSettingsClient)(nil).List), arg0)
}

type MockSecurityContactsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecurityContactsClientMockRecorder
}

type MockSecurityContactsClientMockRecorder struct {
	mock *MockSecurityContactsClient
}

func NewMockSecurityContactsClient(ctrl *gomock.Controller) *MockSecurityContactsClient {
	mock := &MockSecurityContactsClient{ctrl: ctrl}
	mock.recorder = &MockSecurityContactsClientMockRecorder{mock}
	return mock
}

func (m *MockSecurityContactsClient) EXPECT() *MockSecurityContactsClientMockRecorder {
	return m.recorder
}

func (m *MockSecurityContactsClient) List(arg0 context.Context) (security.ContactListPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(security.ContactListPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityContactsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecurityContactsClient)(nil).List), arg0)
}

type MockSecurityPricingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecurityPricingsClientMockRecorder
}

type MockSecurityPricingsClientMockRecorder struct {
	mock *MockSecurityPricingsClient
}

func NewMockSecurityPricingsClient(ctrl *gomock.Controller) *MockSecurityPricingsClient {
	mock := &MockSecurityPricingsClient{ctrl: ctrl}
	mock.recorder = &MockSecurityPricingsClientMockRecorder{mock}
	return mock
}

func (m *MockSecurityPricingsClient) EXPECT() *MockSecurityPricingsClientMockRecorder {
	return m.recorder
}

func (m *MockSecurityPricingsClient) List(arg0 context.Context) (security.PricingList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(security.PricingList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityPricingsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecurityPricingsClient)(nil).List), arg0)
}

type MockSecuritySettingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecuritySettingsClientMockRecorder
}

type MockSecuritySettingsClientMockRecorder struct {
	mock *MockSecuritySettingsClient
}

func NewMockSecuritySettingsClient(ctrl *gomock.Controller) *MockSecuritySettingsClient {
	mock := &MockSecuritySettingsClient{ctrl: ctrl}
	mock.recorder = &MockSecuritySettingsClientMockRecorder{mock}
	return mock
}

func (m *MockSecuritySettingsClient) EXPECT() *MockSecuritySettingsClientMockRecorder {
	return m.recorder
}

func (m *MockSecuritySettingsClient) List(arg0 context.Context) (security.SettingsListPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(security.SettingsListPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecuritySettingsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecuritySettingsClient)(nil).List), arg0)
}

type MockSecurityJitNetworkAccessPoliciesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecurityJitNetworkAccessPoliciesClientMockRecorder
}

type MockSecurityJitNetworkAccessPoliciesClientMockRecorder struct {
	mock *MockSecurityJitNetworkAccessPoliciesClient
}

func NewMockSecurityJitNetworkAccessPoliciesClient(ctrl *gomock.Controller) *MockSecurityJitNetworkAccessPoliciesClient {
	mock := &MockSecurityJitNetworkAccessPoliciesClient{ctrl: ctrl}
	mock.recorder = &MockSecurityJitNetworkAccessPoliciesClientMockRecorder{mock}
	return mock
}

func (m *MockSecurityJitNetworkAccessPoliciesClient) EXPECT() *MockSecurityJitNetworkAccessPoliciesClientMockRecorder {
	return m.recorder
}

func (m *MockSecurityJitNetworkAccessPoliciesClient) List(arg0 context.Context) (security.JitNetworkAccessPoliciesListPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(security.JitNetworkAccessPoliciesListPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityJitNetworkAccessPoliciesClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecurityJitNetworkAccessPoliciesClient)(nil).List), arg0)
}

type MockSecurityAssessmentsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSecurityAssessmentsClientMockRecorder
}

type MockSecurityAssessmentsClientMockRecorder struct {
	mock *MockSecurityAssessmentsClient
}

func NewMockSecurityAssessmentsClient(ctrl *gomock.Controller) *MockSecurityAssessmentsClient {
	mock := &MockSecurityAssessmentsClient{ctrl: ctrl}
	mock.recorder = &MockSecurityAssessmentsClientMockRecorder{mock}
	return mock
}

func (m *MockSecurityAssessmentsClient) EXPECT() *MockSecurityAssessmentsClientMockRecorder {
	return m.recorder
}

func (m *MockSecurityAssessmentsClient) List(arg0 context.Context, arg1 string) (security.AssessmentListPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(security.AssessmentListPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSecurityAssessmentsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSecurityAssessmentsClient)(nil).List), arg0, arg1)
}
