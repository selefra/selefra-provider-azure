package mocks

import (
	context "context"
	reflect "reflect"

	devices "github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2021-07-02/devices"
	gomock "github.com/golang/mock/gomock"
)

type MockIotHubDevicesClient struct {
	ctrl     *gomock.Controller
	recorder *MockIotHubDevicesClientMockRecorder
}

type MockIotHubDevicesClientMockRecorder struct {
	mock *MockIotHubDevicesClient
}

func NewMockIotHubDevicesClient(ctrl *gomock.Controller) *MockIotHubDevicesClient {
	mock := &MockIotHubDevicesClient{ctrl: ctrl}
	mock.recorder = &MockIotHubDevicesClientMockRecorder{mock}
	return mock
}

func (m *MockIotHubDevicesClient) EXPECT() *MockIotHubDevicesClientMockRecorder {
	return m.recorder
}

func (m *MockIotHubDevicesClient) ListBySubscription(arg0 context.Context) (devices.IotHubDescriptionListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscription", arg0)
	ret0, _ := ret[0].(devices.IotHubDescriptionListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockIotHubDevicesClientMockRecorder) ListBySubscription(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscription", reflect.TypeOf((*MockIotHubDevicesClient)(nil).ListBySubscription), arg0)
}
