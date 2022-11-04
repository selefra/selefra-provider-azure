package mocks

import (
	runtime "github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	armsubscriptions "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

type MockSubscriptionsSubscriptionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionsSubscriptionsClientMockRecorder
}

type MockSubscriptionsSubscriptionsClientMockRecorder struct {
	mock *MockSubscriptionsSubscriptionsClient
}

func NewMockSubscriptionsSubscriptionsClient(ctrl *gomock.Controller) *MockSubscriptionsSubscriptionsClient {
	mock := &MockSubscriptionsSubscriptionsClient{ctrl: ctrl}
	mock.recorder = &MockSubscriptionsSubscriptionsClientMockRecorder{mock}
	return mock
}

func (m *MockSubscriptionsSubscriptionsClient) EXPECT() *MockSubscriptionsSubscriptionsClientMockRecorder {
	return m.recorder
}

func (m *MockSubscriptionsSubscriptionsClient) NewListPager(arg0 *armsubscriptions.ClientListOptions) *runtime.Pager[armsubscriptions.ClientListResponse] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0)
	ret0, _ := ret[0].(*runtime.Pager[armsubscriptions.ClientListResponse])
	return ret0
}

func (mr *MockSubscriptionsSubscriptionsClientMockRecorder) NewListPager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockSubscriptionsSubscriptionsClient)(nil).NewListPager), arg0)
}

type MockSubscriptionsTenantsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionsTenantsClientMockRecorder
}

type MockSubscriptionsTenantsClientMockRecorder struct {
	mock *MockSubscriptionsTenantsClient
}

func NewMockSubscriptionsTenantsClient(ctrl *gomock.Controller) *MockSubscriptionsTenantsClient {
	mock := &MockSubscriptionsTenantsClient{ctrl: ctrl}
	mock.recorder = &MockSubscriptionsTenantsClientMockRecorder{mock}
	return mock
}

func (m *MockSubscriptionsTenantsClient) EXPECT() *MockSubscriptionsTenantsClientMockRecorder {
	return m.recorder
}

func (m *MockSubscriptionsTenantsClient) NewListPager(arg0 *armsubscriptions.TenantsClientListOptions) *runtime.Pager[armsubscriptions.TenantsClientListResponse] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListPager", arg0)
	ret0, _ := ret[0].(*runtime.Pager[armsubscriptions.TenantsClientListResponse])
	return ret0
}

func (mr *MockSubscriptionsTenantsClientMockRecorder) NewListPager(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListPager", reflect.TypeOf((*MockSubscriptionsTenantsClient)(nil).NewListPager), arg0)
}

type MockSubscriptionsLocationsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionsLocationsClientMockRecorder
}

type MockSubscriptionsLocationsClientMockRecorder struct {
	mock *MockSubscriptionsLocationsClient
}

func NewMockSubscriptionsLocationsClient(ctrl *gomock.Controller) *MockSubscriptionsLocationsClient {
	mock := &MockSubscriptionsLocationsClient{ctrl: ctrl}
	mock.recorder = &MockSubscriptionsLocationsClientMockRecorder{mock}
	return mock
}

func (m *MockSubscriptionsLocationsClient) EXPECT() *MockSubscriptionsLocationsClientMockRecorder {
	return m.recorder
}

func (m *MockSubscriptionsLocationsClient) NewListLocationsPager(arg0 string, arg1 *armsubscriptions.ClientListLocationsOptions) *runtime.Pager[armsubscriptions.ClientListLocationsResponse] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewListLocationsPager", arg0, arg1)
	ret0, _ := ret[0].(*runtime.Pager[armsubscriptions.ClientListLocationsResponse])
	return ret0
}

func (mr *MockSubscriptionsLocationsClientMockRecorder) NewListLocationsPager(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewListLocationsPager", reflect.TypeOf((*MockSubscriptionsLocationsClient)(nil).NewListLocationsPager), arg0, arg1)
}
