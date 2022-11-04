package mocks

import (
	context "context"
	reflect "reflect"

	redis "github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2020-12-01/redis"
	gomock "github.com/golang/mock/gomock"
)

type MockRedisCachesClient struct {
	ctrl     *gomock.Controller
	recorder *MockRedisCachesClientMockRecorder
}

type MockRedisCachesClientMockRecorder struct {
	mock *MockRedisCachesClient
}

func NewMockRedisCachesClient(ctrl *gomock.Controller) *MockRedisCachesClient {
	mock := &MockRedisCachesClient{ctrl: ctrl}
	mock.recorder = &MockRedisCachesClientMockRecorder{mock}
	return mock
}

func (m *MockRedisCachesClient) EXPECT() *MockRedisCachesClientMockRecorder {
	return m.recorder
}

func (m *MockRedisCachesClient) ListBySubscription(arg0 context.Context) (redis.ListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscription", arg0)
	ret0, _ := ret[0].(redis.ListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockRedisCachesClientMockRecorder) ListBySubscription(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscription", reflect.TypeOf((*MockRedisCachesClient)(nil).ListBySubscription), arg0)
}
