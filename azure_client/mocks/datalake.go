package mocks

import (
	context "context"
	reflect "reflect"

	account "github.com/Azure/azure-sdk-for-go/services/datalake/analytics/mgmt/2016-11-01/account"
	account0 "github.com/Azure/azure-sdk-for-go/services/datalake/store/mgmt/2016-11-01/account"
	gomock "github.com/golang/mock/gomock"
)

type MockDataLakeStoreAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockDataLakeStoreAccountsClientMockRecorder
}

type MockDataLakeStoreAccountsClientMockRecorder struct {
	mock *MockDataLakeStoreAccountsClient
}

func NewMockDataLakeStoreAccountsClient(ctrl *gomock.Controller) *MockDataLakeStoreAccountsClient {
	mock := &MockDataLakeStoreAccountsClient{ctrl: ctrl}
	mock.recorder = &MockDataLakeStoreAccountsClientMockRecorder{mock}
	return mock
}

func (m *MockDataLakeStoreAccountsClient) EXPECT() *MockDataLakeStoreAccountsClientMockRecorder {
	return m.recorder
}

func (m *MockDataLakeStoreAccountsClient) Get(arg0 context.Context, arg1, arg2 string) (account0.DataLakeStoreAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(account0.DataLakeStoreAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDataLakeStoreAccountsClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDataLakeStoreAccountsClient)(nil).Get), arg0, arg1, arg2)
}

func (m *MockDataLakeStoreAccountsClient) List(arg0 context.Context, arg1 string, arg2, arg3 *int32, arg4, arg5 string, arg6 *bool) (account0.DataLakeStoreAccountListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(account0.DataLakeStoreAccountListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDataLakeStoreAccountsClientMockRecorder) List(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDataLakeStoreAccountsClient)(nil).List), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

type MockDataLakeAnalyticsAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockDataLakeAnalyticsAccountsClientMockRecorder
}

type MockDataLakeAnalyticsAccountsClientMockRecorder struct {
	mock *MockDataLakeAnalyticsAccountsClient
}

func NewMockDataLakeAnalyticsAccountsClient(ctrl *gomock.Controller) *MockDataLakeAnalyticsAccountsClient {
	mock := &MockDataLakeAnalyticsAccountsClient{ctrl: ctrl}
	mock.recorder = &MockDataLakeAnalyticsAccountsClientMockRecorder{mock}
	return mock
}

func (m *MockDataLakeAnalyticsAccountsClient) EXPECT() *MockDataLakeAnalyticsAccountsClientMockRecorder {
	return m.recorder
}

func (m *MockDataLakeAnalyticsAccountsClient) Get(arg0 context.Context, arg1, arg2 string) (account.DataLakeAnalyticsAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(account.DataLakeAnalyticsAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDataLakeAnalyticsAccountsClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDataLakeAnalyticsAccountsClient)(nil).Get), arg0, arg1, arg2)
}

func (m *MockDataLakeAnalyticsAccountsClient) List(arg0 context.Context, arg1 string, arg2, arg3 *int32, arg4, arg5 string, arg6 *bool) (account.DataLakeAnalyticsAccountListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(account.DataLakeAnalyticsAccountListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDataLakeAnalyticsAccountsClientMockRecorder) List(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDataLakeAnalyticsAccountsClient)(nil).List), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}
