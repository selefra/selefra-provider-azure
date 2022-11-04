package mocks

import (
	context "context"
	reflect "reflect"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
	gomock "github.com/golang/mock/gomock"
)

type MockComputeDisksClient struct {
	ctrl     *gomock.Controller
	recorder *MockComputeDisksClientMockRecorder
}

type MockComputeDisksClientMockRecorder struct {
	mock *MockComputeDisksClient
}

func NewMockComputeDisksClient(ctrl *gomock.Controller) *MockComputeDisksClient {
	mock := &MockComputeDisksClient{ctrl: ctrl}
	mock.recorder = &MockComputeDisksClientMockRecorder{mock}
	return mock
}

func (m *MockComputeDisksClient) EXPECT() *MockComputeDisksClientMockRecorder {
	return m.recorder
}

func (m *MockComputeDisksClient) List(arg0 context.Context) (compute.DiskListPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(compute.DiskListPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockComputeDisksClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockComputeDisksClient)(nil).List), arg0)
}

type MockComputeVirtualMachinesClient struct {
	ctrl     *gomock.Controller
	recorder *MockComputeVirtualMachinesClientMockRecorder
}

type MockComputeVirtualMachinesClientMockRecorder struct {
	mock *MockComputeVirtualMachinesClient
}

func NewMockComputeVirtualMachinesClient(ctrl *gomock.Controller) *MockComputeVirtualMachinesClient {
	mock := &MockComputeVirtualMachinesClient{ctrl: ctrl}
	mock.recorder = &MockComputeVirtualMachinesClientMockRecorder{mock}
	return mock
}

func (m *MockComputeVirtualMachinesClient) EXPECT() *MockComputeVirtualMachinesClientMockRecorder {
	return m.recorder
}

func (m *MockComputeVirtualMachinesClient) ListAll(arg0 context.Context, arg1 string) (compute.VirtualMachineListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0, arg1)
	ret0, _ := ret[0].(compute.VirtualMachineListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockComputeVirtualMachinesClientMockRecorder) ListAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockComputeVirtualMachinesClient)(nil).ListAll), arg0, arg1)
}

type MockComputeVirtualMachineExtensionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockComputeVirtualMachineExtensionsClientMockRecorder
}

type MockComputeVirtualMachineExtensionsClientMockRecorder struct {
	mock *MockComputeVirtualMachineExtensionsClient
}

func NewMockComputeVirtualMachineExtensionsClient(ctrl *gomock.Controller) *MockComputeVirtualMachineExtensionsClient {
	mock := &MockComputeVirtualMachineExtensionsClient{ctrl: ctrl}
	mock.recorder = &MockComputeVirtualMachineExtensionsClientMockRecorder{mock}
	return mock
}

func (m *MockComputeVirtualMachineExtensionsClient) EXPECT() *MockComputeVirtualMachineExtensionsClientMockRecorder {
	return m.recorder
}

func (m *MockComputeVirtualMachineExtensionsClient) List(arg0 context.Context, arg1, arg2, arg3 string) (compute.VirtualMachineExtensionsListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(compute.VirtualMachineExtensionsListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockComputeVirtualMachineExtensionsClientMockRecorder) List(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockComputeVirtualMachineExtensionsClient)(nil).List), arg0, arg1, arg2, arg3)
}

type MockComputeVirtualMachineScaleSetsClient struct {
	ctrl     *gomock.Controller
	recorder *MockComputeVirtualMachineScaleSetsClientMockRecorder
}

type MockComputeVirtualMachineScaleSetsClientMockRecorder struct {
	mock *MockComputeVirtualMachineScaleSetsClient
}

func NewMockComputeVirtualMachineScaleSetsClient(ctrl *gomock.Controller) *MockComputeVirtualMachineScaleSetsClient {
	mock := &MockComputeVirtualMachineScaleSetsClient{ctrl: ctrl}
	mock.recorder = &MockComputeVirtualMachineScaleSetsClientMockRecorder{mock}
	return mock
}

func (m *MockComputeVirtualMachineScaleSetsClient) EXPECT() *MockComputeVirtualMachineScaleSetsClientMockRecorder {
	return m.recorder
}

func (m *MockComputeVirtualMachineScaleSetsClient) ListAll(arg0 context.Context) (compute.VirtualMachineScaleSetListWithLinkResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetListWithLinkResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockComputeVirtualMachineScaleSetsClientMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockComputeVirtualMachineScaleSetsClient)(nil).ListAll), arg0)
}

type MockComputeInstanceViewsClient struct {
	ctrl     *gomock.Controller
	recorder *MockComputeInstanceViewsClientMockRecorder
}

type MockComputeInstanceViewsClientMockRecorder struct {
	mock *MockComputeInstanceViewsClient
}

func NewMockComputeInstanceViewsClient(ctrl *gomock.Controller) *MockComputeInstanceViewsClient {
	mock := &MockComputeInstanceViewsClient{ctrl: ctrl}
	mock.recorder = &MockComputeInstanceViewsClientMockRecorder{mock}
	return mock
}

func (m *MockComputeInstanceViewsClient) EXPECT() *MockComputeInstanceViewsClientMockRecorder {
	return m.recorder
}

func (m *MockComputeInstanceViewsClient) InstanceView(arg0 context.Context, arg1, arg2 string) (compute.VirtualMachineInstanceView, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InstanceView", arg0, arg1, arg2)
	ret0, _ := ret[0].(compute.VirtualMachineInstanceView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockComputeInstanceViewsClientMockRecorder) InstanceView(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceView", reflect.TypeOf((*MockComputeInstanceViewsClient)(nil).InstanceView), arg0, arg1, arg2)
}
