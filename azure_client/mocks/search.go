package mocks

import (
	context "context"
	reflect "reflect"

	search "github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-08-01/search"
	uuid "github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
)

type MockSearchServicesClient struct {
	ctrl     *gomock.Controller
	recorder *MockSearchServicesClientMockRecorder
}

type MockSearchServicesClientMockRecorder struct {
	mock *MockSearchServicesClient
}

func NewMockSearchServicesClient(ctrl *gomock.Controller) *MockSearchServicesClient {
	mock := &MockSearchServicesClient{ctrl: ctrl}
	mock.recorder = &MockSearchServicesClientMockRecorder{mock}
	return mock
}

func (m *MockSearchServicesClient) EXPECT() *MockSearchServicesClientMockRecorder {
	return m.recorder
}

func (m *MockSearchServicesClient) ListBySubscription(arg0 context.Context, arg1 *uuid.UUID) (search.ServiceListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscription", arg0, arg1)
	ret0, _ := ret[0].(search.ServiceListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockSearchServicesClientMockRecorder) ListBySubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscription", reflect.TypeOf((*MockSearchServicesClient)(nil).ListBySubscription), arg0, arg1)
}
