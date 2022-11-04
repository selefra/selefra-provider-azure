package mocks

import (
	context "context"
	reflect "reflect"

	keyvault "github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2019-09-01/keyvault"
	keyvault0 "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	keyvault1 "github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"
	gomock "github.com/golang/mock/gomock"
)

type MockKeyVaultVaultsClient struct {
	ctrl     *gomock.Controller
	recorder *MockKeyVaultVaultsClientMockRecorder
}

type MockKeyVaultVaultsClientMockRecorder struct {
	mock *MockKeyVaultVaultsClient
}

func NewMockKeyVaultVaultsClient(ctrl *gomock.Controller) *MockKeyVaultVaultsClient {
	mock := &MockKeyVaultVaultsClient{ctrl: ctrl}
	mock.recorder = &MockKeyVaultVaultsClientMockRecorder{mock}
	return mock
}

func (m *MockKeyVaultVaultsClient) EXPECT() *MockKeyVaultVaultsClientMockRecorder {
	return m.recorder
}

func (m *MockKeyVaultVaultsClient) ListBySubscription(arg0 context.Context, arg1 *int32) (keyvault.VaultListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscription", arg0, arg1)
	ret0, _ := ret[0].(keyvault.VaultListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKeyVaultVaultsClientMockRecorder) ListBySubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscription", reflect.TypeOf((*MockKeyVaultVaultsClient)(nil).ListBySubscription), arg0, arg1)
}

type MockKeyVaultManagedHsmsClient struct {
	ctrl     *gomock.Controller
	recorder *MockKeyVaultManagedHsmsClientMockRecorder
}

type MockKeyVaultManagedHsmsClientMockRecorder struct {
	mock *MockKeyVaultManagedHsmsClient
}

func NewMockKeyVaultManagedHsmsClient(ctrl *gomock.Controller) *MockKeyVaultManagedHsmsClient {
	mock := &MockKeyVaultManagedHsmsClient{ctrl: ctrl}
	mock.recorder = &MockKeyVaultManagedHsmsClientMockRecorder{mock}
	return mock
}

func (m *MockKeyVaultManagedHsmsClient) EXPECT() *MockKeyVaultManagedHsmsClientMockRecorder {
	return m.recorder
}

func (m *MockKeyVaultManagedHsmsClient) ListBySubscription(arg0 context.Context, arg1 *int32) (keyvault1.ManagedHsmListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscription", arg0, arg1)
	ret0, _ := ret[0].(keyvault1.ManagedHsmListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKeyVaultManagedHsmsClientMockRecorder) ListBySubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscription", reflect.TypeOf((*MockKeyVaultManagedHsmsClient)(nil).ListBySubscription), arg0, arg1)
}

type MockKeyVaultKeysClient struct {
	ctrl     *gomock.Controller
	recorder *MockKeyVaultKeysClientMockRecorder
}

type MockKeyVaultKeysClientMockRecorder struct {
	mock *MockKeyVaultKeysClient
}

func NewMockKeyVaultKeysClient(ctrl *gomock.Controller) *MockKeyVaultKeysClient {
	mock := &MockKeyVaultKeysClient{ctrl: ctrl}
	mock.recorder = &MockKeyVaultKeysClientMockRecorder{mock}
	return mock
}

func (m *MockKeyVaultKeysClient) EXPECT() *MockKeyVaultKeysClientMockRecorder {
	return m.recorder
}

func (m *MockKeyVaultKeysClient) GetKeys(arg0 context.Context, arg1 string, arg2 *int32) (keyvault0.KeyListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKeys", arg0, arg1, arg2)
	ret0, _ := ret[0].(keyvault0.KeyListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKeyVaultKeysClientMockRecorder) GetKeys(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKeys", reflect.TypeOf((*MockKeyVaultKeysClient)(nil).GetKeys), arg0, arg1, arg2)
}

type MockKeyVaultSecretsClient struct {
	ctrl     *gomock.Controller
	recorder *MockKeyVaultSecretsClientMockRecorder
}

type MockKeyVaultSecretsClientMockRecorder struct {
	mock *MockKeyVaultSecretsClient
}

func NewMockKeyVaultSecretsClient(ctrl *gomock.Controller) *MockKeyVaultSecretsClient {
	mock := &MockKeyVaultSecretsClient{ctrl: ctrl}
	mock.recorder = &MockKeyVaultSecretsClientMockRecorder{mock}
	return mock
}

func (m *MockKeyVaultSecretsClient) EXPECT() *MockKeyVaultSecretsClientMockRecorder {
	return m.recorder
}

func (m *MockKeyVaultSecretsClient) GetSecrets(arg0 context.Context, arg1 string, arg2 *int32) (keyvault0.SecretListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecrets", arg0, arg1, arg2)
	ret0, _ := ret[0].(keyvault0.SecretListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockKeyVaultSecretsClientMockRecorder) GetSecrets(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecrets", reflect.TypeOf((*MockKeyVaultSecretsClient)(nil).GetSecrets), arg0, arg1, arg2)
}
