package mocks

import (
	context "context"
	reflect "reflect"

	frontdoor "github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor"
	gomock "github.com/golang/mock/gomock"
)

type MockFrontDoorDoorsClient struct {
	ctrl     *gomock.Controller
	recorder *MockFrontDoorDoorsClientMockRecorder
}

type MockFrontDoorDoorsClientMockRecorder struct {
	mock *MockFrontDoorDoorsClient
}

func NewMockFrontDoorDoorsClient(ctrl *gomock.Controller) *MockFrontDoorDoorsClient {
	mock := &MockFrontDoorDoorsClient{ctrl: ctrl}
	mock.recorder = &MockFrontDoorDoorsClientMockRecorder{mock}
	return mock
}

func (m *MockFrontDoorDoorsClient) EXPECT() *MockFrontDoorDoorsClientMockRecorder {
	return m.recorder
}

func (m *MockFrontDoorDoorsClient) List(arg0 context.Context) (frontdoor.ListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(frontdoor.ListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockFrontDoorDoorsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFrontDoorDoorsClient)(nil).List), arg0)
}
