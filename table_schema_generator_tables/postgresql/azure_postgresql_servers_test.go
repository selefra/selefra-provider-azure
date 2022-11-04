package postgresql

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2020-01-01/postgresql"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createConfigurationsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockPostgreSQLConfigurationsClient(ctrl)
	s := services.Services{
		PostgreSQL: services.PostgreSQLClient{
			Configurations: mockClient,
		},
	}

	data := postgresql.Configuration{}
	require.Nil(t, faker.FakeObject(&data))

	result := postgresql.ConfigurationListResult{Value: &[]postgresql.Configuration{data}}

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createFirewallRulesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockPostgreSQLFirewallRulesClient(ctrl)
	s := services.Services{
		PostgreSQL: services.PostgreSQLClient{
			FirewallRules: mockClient,
		},
	}

	data := postgresql.FirewallRule{}
	require.Nil(t, faker.FakeObject(&data))

	result := postgresql.FirewallRuleListResult{Value: &[]postgresql.FirewallRule{data}}

	mockClient.EXPECT().ListByServer(gomock.Any(), "test", "test").Return(result, nil)
	return s
}

func createServersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockPostgreSQLServersClient(ctrl)
	s := services.Services{
		PostgreSQL: services.PostgreSQLClient{
			Servers:        mockClient,
			Configurations: createConfigurationsMock(t, ctrl).PostgreSQL.Configurations,
			FirewallRules:  createFirewallRulesMock(t, ctrl).PostgreSQL.FirewallRules,
		},
	}

	data := postgresql.Server{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := postgresql.ServerListResult{Value: &[]postgresql.Server{data}}

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestPostgreSQLServers(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzurePostgresqlServersGenerator{}), createServersMock, azure_client.TestOptions{})
}
