package mocks

import (
	context "context"
	reflect "reflect"

	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	gomock "github.com/golang/mock/gomock"
)

type MockNetworkExpressRouteCircuitsClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkExpressRouteCircuitsClientMockRecorder
}

type MockNetworkExpressRouteCircuitsClientMockRecorder struct {
	mock *MockNetworkExpressRouteCircuitsClient
}

func NewMockNetworkExpressRouteCircuitsClient(ctrl *gomock.Controller) *MockNetworkExpressRouteCircuitsClient {
	mock := &MockNetworkExpressRouteCircuitsClient{ctrl: ctrl}
	mock.recorder = &MockNetworkExpressRouteCircuitsClientMockRecorder{mock}
	return mock
}

func (m *MockNetworkExpressRouteCircuitsClient) EXPECT() *MockNetworkExpressRouteCircuitsClientMockRecorder {
	return m.recorder
}

func (m *MockNetworkExpressRouteCircuitsClient) ListAll(arg0 context.Context) (network.ExpressRouteCircuitListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.ExpressRouteCircuitListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkExpressRouteCircuitsClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockNetworkExpressRouteCircuitsClient)(nil).ListAll), arg0)
}

type MockNetworkExpressRouteGatewaysClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkExpressRouteGatewaysClientMockRecorder
}

type MockNetworkExpressRouteGatewaysClientMockRecorder struct {
	mock *MockNetworkExpressRouteGatewaysClient
}

func NewMockNetworkExpressRouteGatewaysClient(ctrl *gomock.Controller) *MockNetworkExpressRouteGatewaysClient {
	mock := &MockNetworkExpressRouteGatewaysClient{ctrl: ctrl}
	mock.recorder = &MockNetworkExpressRouteGatewaysClientMockRecorder{mock}
	return mock
}

func (m *MockNetworkExpressRouteGatewaysClient) EXPECT() *MockNetworkExpressRouteGatewaysClientMockRecorder {
	return m.recorder
}

func (m *MockNetworkExpressRouteGatewaysClient) ListBySubscription(arg0 context.Context) (network.ExpressRouteGatewayList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBySubscription", arg0)
	ret0, _ := ret[0].(network.ExpressRouteGatewayList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkExpressRouteGatewaysClientMockRecorder) ListBySubscription(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBySubscription", reflect.TypeOf((*MockNetworkExpressRouteGatewaysClient)(nil).ListBySubscription), arg0)
}

type MockNetworkExpressRoutePortsClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkExpressRoutePortsClientMockRecorder
}

type MockNetworkExpressRoutePortsClientMockRecorder struct {
	mock *MockNetworkExpressRoutePortsClient
}

func NewMockNetworkExpressRoutePortsClient(ctrl *gomock.Controller) *MockNetworkExpressRoutePortsClient {
	mock := &MockNetworkExpressRoutePortsClient{ctrl: ctrl}
	mock.recorder = &MockNetworkExpressRoutePortsClientMockRecorder{mock}
	return mock
}

func (m *MockNetworkExpressRoutePortsClient) EXPECT() *MockNetworkExpressRoutePortsClientMockRecorder {
	return m.recorder
}

func (m *MockNetworkExpressRoutePortsClient) List(arg0 context.Context) (network.ExpressRoutePortListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(network.ExpressRoutePortListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkExpressRoutePortsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNetworkExpressRoutePortsClient)(nil).List), arg0)
}

type MockNetworkInterfacesClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkInterfacesClientMockRecorder
}

type MockNetworkInterfacesClientMockRecorder struct {
	mock *MockNetworkInterfacesClient
}

func NewMockNetworkInterfacesClient(ctrl *gomock.Controller) *MockNetworkInterfacesClient {
	mock := &MockNetworkInterfacesClient{ctrl: ctrl}
	mock.recorder = &MockNetworkInterfacesClientMockRecorder{mock}
	return mock
}

func (m *MockNetworkInterfacesClient) EXPECT() *MockNetworkInterfacesClientMockRecorder {
	return m.recorder
}

func (m *MockNetworkInterfacesClient) ListAll(arg0 context.Context) (network.InterfaceListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.InterfaceListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkInterfacesClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockNetworkInterfacesClient)(nil).ListAll), arg0)
}

type MockNetworkPublicIPAddressesClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkPublicIPAddressesClientMockRecorder
}

type MockNetworkPublicIPAddressesClientMockRecorder struct {
	mock *MockNetworkPublicIPAddressesClient
}

func NewMockNetworkPublicIPAddressesClient(ctrl *gomock.Controller) *MockNetworkPublicIPAddressesClient {
	mock := &MockNetworkPublicIPAddressesClient{ctrl: ctrl}
	mock.recorder = &MockNetworkPublicIPAddressesClientMockRecorder{mock}
	return mock
}

func (m *MockNetworkPublicIPAddressesClient) EXPECT() *MockNetworkPublicIPAddressesClientMockRecorder {
	return m.recorder
}

func (m *MockNetworkPublicIPAddressesClient) ListAll(arg0 context.Context) (network.PublicIPAddressListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.PublicIPAddressListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkPublicIPAddressesClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockNetworkPublicIPAddressesClient)(nil).ListAll), arg0)
}

type MockNetworkRouteFiltersClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkRouteFiltersClientMockRecorder
}

type MockNetworkRouteFiltersClientMockRecorder struct {
	mock *MockNetworkRouteFiltersClient
}

func NewMockNetworkRouteFiltersClient(ctrl *gomock.Controller) *MockNetworkRouteFiltersClient {
	mock := &MockNetworkRouteFiltersClient{ctrl: ctrl}
	mock.recorder = &MockNetworkRouteFiltersClientMockRecorder{mock}
	return mock
}

func (m *MockNetworkRouteFiltersClient) EXPECT() *MockNetworkRouteFiltersClientMockRecorder {
	return m.recorder
}

func (m *MockNetworkRouteFiltersClient) List(arg0 context.Context) (network.RouteFilterListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(network.RouteFilterListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkRouteFiltersClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNetworkRouteFiltersClient)(nil).List), arg0)
}

type MockNetworkRouteTablesClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkRouteTablesClientMockRecorder
}

type MockNetworkRouteTablesClientMockRecorder struct {
	mock *MockNetworkRouteTablesClient
}

func NewMockNetworkRouteTablesClient(ctrl *gomock.Controller) *MockNetworkRouteTablesClient {
	mock := &MockNetworkRouteTablesClient{ctrl: ctrl}
	mock.recorder = &MockNetworkRouteTablesClientMockRecorder{mock}
	return mock
}

func (m *MockNetworkRouteTablesClient) EXPECT() *MockNetworkRouteTablesClientMockRecorder {
	return m.recorder
}

func (m *MockNetworkRouteTablesClient) ListAll(arg0 context.Context) (network.RouteTableListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.RouteTableListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkRouteTablesClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockNetworkRouteTablesClient)(nil).ListAll), arg0)
}

type MockNetworkSecurityGroupsClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkSecurityGroupsClientMockRecorder
}

type MockNetworkSecurityGroupsClientMockRecorder struct {
	mock *MockNetworkSecurityGroupsClient
}

func NewMockNetworkSecurityGroupsClient(ctrl *gomock.Controller) *MockNetworkSecurityGroupsClient {
	mock := &MockNetworkSecurityGroupsClient{ctrl: ctrl}
	mock.recorder = &MockNetworkSecurityGroupsClientMockRecorder{mock}
	return mock
}

func (m *MockNetworkSecurityGroupsClient) EXPECT() *MockNetworkSecurityGroupsClientMockRecorder {
	return m.recorder
}

func (m *MockNetworkSecurityGroupsClient) ListAll(arg0 context.Context) (network.SecurityGroupListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.SecurityGroupListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkSecurityGroupsClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockNetworkSecurityGroupsClient)(nil).ListAll), arg0)
}

type MockNetworkVirtualNetworkGatewaysClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkVirtualNetworkGatewaysClientMockRecorder
}

type MockNetworkVirtualNetworkGatewaysClientMockRecorder struct {
	mock *MockNetworkVirtualNetworkGatewaysClient
}

func NewMockNetworkVirtualNetworkGatewaysClient(ctrl *gomock.Controller) *MockNetworkVirtualNetworkGatewaysClient {
	mock := &MockNetworkVirtualNetworkGatewaysClient{ctrl: ctrl}
	mock.recorder = &MockNetworkVirtualNetworkGatewaysClientMockRecorder{mock}
	return mock
}

func (m *MockNetworkVirtualNetworkGatewaysClient) EXPECT() *MockNetworkVirtualNetworkGatewaysClientMockRecorder {
	return m.recorder
}

func (m *MockNetworkVirtualNetworkGatewaysClient) List(arg0 context.Context, arg1 string) (network.VirtualNetworkGatewayListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(network.VirtualNetworkGatewayListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkVirtualNetworkGatewaysClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNetworkVirtualNetworkGatewaysClient)(nil).List), arg0, arg1)
}

type MockNetworkVirtualNetworkGatewayConnectionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkVirtualNetworkGatewayConnectionsClientMockRecorder
}

type MockNetworkVirtualNetworkGatewayConnectionsClientMockRecorder struct {
	mock *MockNetworkVirtualNetworkGatewayConnectionsClient
}

func NewMockNetworkVirtualNetworkGatewayConnectionsClient(ctrl *gomock.Controller) *MockNetworkVirtualNetworkGatewayConnectionsClient {
	mock := &MockNetworkVirtualNetworkGatewayConnectionsClient{ctrl: ctrl}
	mock.recorder = &MockNetworkVirtualNetworkGatewayConnectionsClientMockRecorder{mock}
	return mock
}

func (m *MockNetworkVirtualNetworkGatewayConnectionsClient) EXPECT() *MockNetworkVirtualNetworkGatewayConnectionsClientMockRecorder {
	return m.recorder
}

func (m *MockNetworkVirtualNetworkGatewayConnectionsClient) ListConnections(arg0 context.Context, arg1, arg2 string) (network.VirtualNetworkGatewayListConnectionsResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListConnections", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.VirtualNetworkGatewayListConnectionsResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkVirtualNetworkGatewayConnectionsClientMockRecorder) ListConnections(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListConnections", reflect.TypeOf((*MockNetworkVirtualNetworkGatewayConnectionsClient)(nil).ListConnections), arg0, arg1, arg2)
}

type MockNetworkVirtualNetworksClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkVirtualNetworksClientMockRecorder
}

type MockNetworkVirtualNetworksClientMockRecorder struct {
	mock *MockNetworkVirtualNetworksClient
}

func NewMockNetworkVirtualNetworksClient(ctrl *gomock.Controller) *MockNetworkVirtualNetworksClient {
	mock := &MockNetworkVirtualNetworksClient{ctrl: ctrl}
	mock.recorder = &MockNetworkVirtualNetworksClientMockRecorder{mock}
	return mock
}

func (m *MockNetworkVirtualNetworksClient) EXPECT() *MockNetworkVirtualNetworksClientMockRecorder {
	return m.recorder
}

func (m *MockNetworkVirtualNetworksClient) ListAll(arg0 context.Context) (network.VirtualNetworkListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.VirtualNetworkListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkVirtualNetworksClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockNetworkVirtualNetworksClient)(nil).ListAll), arg0)
}

type MockNetworkWatchersClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkWatchersClientMockRecorder
}

type MockNetworkWatchersClientMockRecorder struct {
	mock *MockNetworkWatchersClient
}

func NewMockNetworkWatchersClient(ctrl *gomock.Controller) *MockNetworkWatchersClient {
	mock := &MockNetworkWatchersClient{ctrl: ctrl}
	mock.recorder = &MockNetworkWatchersClientMockRecorder{mock}
	return mock
}

func (m *MockNetworkWatchersClient) EXPECT() *MockNetworkWatchersClientMockRecorder {
	return m.recorder
}

func (m *MockNetworkWatchersClient) ListAll(arg0 context.Context) (network.WatcherListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.WatcherListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkWatchersClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockNetworkWatchersClient)(nil).ListAll), arg0)
}

type MockNetworkFlowLogsClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkFlowLogsClientMockRecorder
}

type MockNetworkFlowLogsClientMockRecorder struct {
	mock *MockNetworkFlowLogsClient
}

func NewMockNetworkFlowLogsClient(ctrl *gomock.Controller) *MockNetworkFlowLogsClient {
	mock := &MockNetworkFlowLogsClient{ctrl: ctrl}
	mock.recorder = &MockNetworkFlowLogsClientMockRecorder{mock}
	return mock
}

func (m *MockNetworkFlowLogsClient) EXPECT() *MockNetworkFlowLogsClientMockRecorder {
	return m.recorder
}

func (m *MockNetworkFlowLogsClient) List(arg0 context.Context, arg1, arg2 string) (network.FlowLogListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.FlowLogListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockNetworkFlowLogsClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNetworkFlowLogsClient)(nil).List), arg0, arg1, arg2)
}
