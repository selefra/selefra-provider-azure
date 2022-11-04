package mocks

import (
	context "context"
	reflect "reflect"

	web "github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	gomock "github.com/golang/mock/gomock"

	services "github.com/selefra/selefra-provider-azure/azure_client/services"
)

type MockWebAppsClient struct {
	ctrl     *gomock.Controller
	recorder *MockWebAppsClientMockRecorder
}

type MockWebAppsClientMockRecorder struct {
	mock *MockWebAppsClient
}

func NewMockWebAppsClient(ctrl *gomock.Controller) *MockWebAppsClient {
	mock := &MockWebAppsClient{ctrl: ctrl}
	mock.recorder = &MockWebAppsClientMockRecorder{mock}
	return mock
}

func (m *MockWebAppsClient) EXPECT() *MockWebAppsClientMockRecorder {
	return m.recorder
}

func (m *MockWebAppsClient) List(arg0 context.Context) (web.AppCollectionPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(web.AppCollectionPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWebAppsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockWebAppsClient)(nil).List), arg0)
}

type MockWebSiteAuthSettingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockWebSiteAuthSettingsClientMockRecorder
}

type MockWebSiteAuthSettingsClientMockRecorder struct {
	mock *MockWebSiteAuthSettingsClient
}

func NewMockWebSiteAuthSettingsClient(ctrl *gomock.Controller) *MockWebSiteAuthSettingsClient {
	mock := &MockWebSiteAuthSettingsClient{ctrl: ctrl}
	mock.recorder = &MockWebSiteAuthSettingsClientMockRecorder{mock}
	return mock
}

func (m *MockWebSiteAuthSettingsClient) EXPECT() *MockWebSiteAuthSettingsClientMockRecorder {
	return m.recorder
}

func (m *MockWebSiteAuthSettingsClient) GetAuthSettings(arg0 context.Context, arg1, arg2 string) (web.SiteAuthSettings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthSettings", arg0, arg1, arg2)
	ret0, _ := ret[0].(web.SiteAuthSettings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWebSiteAuthSettingsClientMockRecorder) GetAuthSettings(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthSettings", reflect.TypeOf((*MockWebSiteAuthSettingsClient)(nil).GetAuthSettings), arg0, arg1, arg2)
}

type MockWebVnetConnectionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockWebVnetConnectionsClientMockRecorder
}

type MockWebVnetConnectionsClientMockRecorder struct {
	mock *MockWebVnetConnectionsClient
}

func NewMockWebVnetConnectionsClient(ctrl *gomock.Controller) *MockWebVnetConnectionsClient {
	mock := &MockWebVnetConnectionsClient{ctrl: ctrl}
	mock.recorder = &MockWebVnetConnectionsClientMockRecorder{mock}
	return mock
}

func (m *MockWebVnetConnectionsClient) EXPECT() *MockWebVnetConnectionsClientMockRecorder {
	return m.recorder
}

func (m *MockWebVnetConnectionsClient) GetVnetConnection(arg0 context.Context, arg1, arg2, arg3 string) (web.VnetInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVnetConnection", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(web.VnetInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWebVnetConnectionsClientMockRecorder) GetVnetConnection(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVnetConnection", reflect.TypeOf((*MockWebVnetConnectionsClient)(nil).GetVnetConnection), arg0, arg1, arg2, arg3)
}

type MockWebPublishingProfilesClient struct {
	ctrl     *gomock.Controller
	recorder *MockWebPublishingProfilesClientMockRecorder
}

type MockWebPublishingProfilesClientMockRecorder struct {
	mock *MockWebPublishingProfilesClient
}

func NewMockWebPublishingProfilesClient(ctrl *gomock.Controller) *MockWebPublishingProfilesClient {
	mock := &MockWebPublishingProfilesClient{ctrl: ctrl}
	mock.recorder = &MockWebPublishingProfilesClientMockRecorder{mock}
	return mock
}

func (m *MockWebPublishingProfilesClient) EXPECT() *MockWebPublishingProfilesClientMockRecorder {
	return m.recorder
}

func (m *MockWebPublishingProfilesClient) ListPublishingProfiles(arg0 context.Context, arg1, arg2 string) (services.PublishingProfiles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPublishingProfiles", arg0, arg1, arg2)
	ret0, _ := ret[0].(services.PublishingProfiles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockWebPublishingProfilesClientMockRecorder) ListPublishingProfiles(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPublishingProfiles", reflect.TypeOf((*MockWebPublishingProfilesClient)(nil).ListPublishingProfiles), arg0, arg1, arg2)
}
