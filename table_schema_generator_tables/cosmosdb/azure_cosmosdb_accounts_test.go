package cosmosdb

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createSQLDatabasesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCosmosDBSQLDatabasesClient(ctrl)
	s := services.Services{
		CosmosDB: services.CosmosDBClient{
			SQLDatabases: mockClient,
		},
	}

	data := documentdb.SQLDatabaseGetResults{}
	require.Nil(t, faker.FakeObject(&data))

	result := documentdb.SQLDatabaseListResult{Value: &[]documentdb.SQLDatabaseGetResults{data}}

	mockClient.EXPECT().ListSQLDatabases(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createMongoDBDatabasesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCosmosDBMongoDBDatabasesClient(ctrl)
	s := services.Services{
		CosmosDB: services.CosmosDBClient{
			MongoDBDatabases: mockClient,
		},
	}

	data := documentdb.MongoDBDatabaseGetResults{}
	require.Nil(t, faker.FakeObject(&data))

	result := documentdb.MongoDBDatabaseListResult{Value: &[]documentdb.MongoDBDatabaseGetResults{data}}

	mockClient.EXPECT().ListMongoDBDatabases(gomock.Any(), "test", "test").Return(result, nil)
	return s
}

func createAccountsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockCosmosDBAccountsClient(ctrl)
	s := services.Services{
		CosmosDB: services.CosmosDBClient{
			Accounts:         mockClient,
			MongoDBDatabases: createMongoDBDatabasesMock(t, ctrl).CosmosDB.MongoDBDatabases,
			SQLDatabases:     createSQLDatabasesMock(t, ctrl).CosmosDB.SQLDatabases,
		},
	}

	data := documentdb.DatabaseAccountGetResults{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := documentdb.DatabaseAccountsListResult{Value: &[]documentdb.DatabaseAccountGetResults{data}}

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestCosmosDBAccounts(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureCosmosdbAccountsGenerator{}), createAccountsMock, azure_client.TestOptions{})
}
