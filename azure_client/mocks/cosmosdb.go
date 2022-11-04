package mocks

import (
	context "context"
	reflect "reflect"

	documentdb "github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb"
	gomock "github.com/golang/mock/gomock"
)

type MockCosmosDBAccountsClient struct {
	ctrl     *gomock.Controller
	recorder *MockCosmosDBAccountsClientMockRecorder
}

type MockCosmosDBAccountsClientMockRecorder struct {
	mock *MockCosmosDBAccountsClient
}

func NewMockCosmosDBAccountsClient(ctrl *gomock.Controller) *MockCosmosDBAccountsClient {
	mock := &MockCosmosDBAccountsClient{ctrl: ctrl}
	mock.recorder = &MockCosmosDBAccountsClientMockRecorder{mock}
	return mock
}

func (m *MockCosmosDBAccountsClient) EXPECT() *MockCosmosDBAccountsClientMockRecorder {
	return m.recorder
}

func (m *MockCosmosDBAccountsClient) List(arg0 context.Context) (documentdb.DatabaseAccountsListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(documentdb.DatabaseAccountsListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCosmosDBAccountsClientMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCosmosDBAccountsClient)(nil).List), arg0)
}

type MockCosmosDBSQLDatabasesClient struct {
	ctrl     *gomock.Controller
	recorder *MockCosmosDBSQLDatabasesClientMockRecorder
}

type MockCosmosDBSQLDatabasesClientMockRecorder struct {
	mock *MockCosmosDBSQLDatabasesClient
}

func NewMockCosmosDBSQLDatabasesClient(ctrl *gomock.Controller) *MockCosmosDBSQLDatabasesClient {
	mock := &MockCosmosDBSQLDatabasesClient{ctrl: ctrl}
	mock.recorder = &MockCosmosDBSQLDatabasesClientMockRecorder{mock}
	return mock
}

func (m *MockCosmosDBSQLDatabasesClient) EXPECT() *MockCosmosDBSQLDatabasesClientMockRecorder {
	return m.recorder
}

func (m *MockCosmosDBSQLDatabasesClient) ListSQLDatabases(arg0 context.Context, arg1, arg2 string) (documentdb.SQLDatabaseListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSQLDatabases", arg0, arg1, arg2)
	ret0, _ := ret[0].(documentdb.SQLDatabaseListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCosmosDBSQLDatabasesClientMockRecorder) ListSQLDatabases(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSQLDatabases", reflect.TypeOf((*MockCosmosDBSQLDatabasesClient)(nil).ListSQLDatabases), arg0, arg1, arg2)
}

type MockCosmosDBMongoDBDatabasesClient struct {
	ctrl     *gomock.Controller
	recorder *MockCosmosDBMongoDBDatabasesClientMockRecorder
}

type MockCosmosDBMongoDBDatabasesClientMockRecorder struct {
	mock *MockCosmosDBMongoDBDatabasesClient
}

func NewMockCosmosDBMongoDBDatabasesClient(ctrl *gomock.Controller) *MockCosmosDBMongoDBDatabasesClient {
	mock := &MockCosmosDBMongoDBDatabasesClient{ctrl: ctrl}
	mock.recorder = &MockCosmosDBMongoDBDatabasesClientMockRecorder{mock}
	return mock
}

func (m *MockCosmosDBMongoDBDatabasesClient) EXPECT() *MockCosmosDBMongoDBDatabasesClientMockRecorder {
	return m.recorder
}

func (m *MockCosmosDBMongoDBDatabasesClient) ListMongoDBDatabases(arg0 context.Context, arg1, arg2 string) (documentdb.MongoDBDatabaseListResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMongoDBDatabases", arg0, arg1, arg2)
	ret0, _ := ret[0].(documentdb.MongoDBDatabaseListResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockCosmosDBMongoDBDatabasesClientMockRecorder) ListMongoDBDatabases(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMongoDBDatabases", reflect.TypeOf((*MockCosmosDBMongoDBDatabasesClient)(nil).ListMongoDBDatabases), arg0, arg1, arg2)
}
