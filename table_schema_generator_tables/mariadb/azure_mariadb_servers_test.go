package mariadb

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/mariadb/mgmt/2020-01-01/mariadb"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createConfigurationsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMariaDBConfigurationsClient(ctrl)
	s := services.Services{
		MariaDB: services.MariaDBClient{
			Configurations: mockClient,
		},
	}

	data := mariadb.Configuration{}
	require.Nil(t, faker.FakeObject(&data))

	result := mariadb.ConfigurationListResult{Value: &[]mariadb.Configuration{data}}

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}

func createServersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMariaDBServersClient(ctrl)
	s := services.Services{
		MariaDB: services.MariaDBClient{
			Servers:        mockClient,
			Configurations: createConfigurationsMock(t, ctrl).MariaDB.Configurations,
		},
	}

	data := mariadb.Server{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := mariadb.ServerListResult{Value: &[]mariadb.Server{data}}

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestMariaDBServers(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureMariadbServersGenerator{}), createServersMock, azure_client.TestOptions{})
}
