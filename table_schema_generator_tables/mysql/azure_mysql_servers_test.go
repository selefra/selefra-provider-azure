package mysql

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createConfigurationsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMySQLConfigurationsClient(ctrl)
	s := services.Services{
		MySQL: services.MySQLClient{
			Configurations: mockClient,
		},
	}

	data := mysql.Configuration{}
	require.Nil(t, faker.FakeObject(&data))

	result := mysql.ConfigurationListResult{Value: &[]mysql.Configuration{data}}

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}

func createServersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMySQLServersClient(ctrl)
	s := services.Services{
		MySQL: services.MySQLClient{
			Servers:        mockClient,
			Configurations: createConfigurationsMock(t, ctrl).MySQL.Configurations,
		},
	}

	data := mysql.Server{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := mysql.ServerListResult{Value: &[]mysql.Server{data}}

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestMySQLServers(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureMysqlServersGenerator{}), createServersMock, azure_client.TestOptions{})
}
