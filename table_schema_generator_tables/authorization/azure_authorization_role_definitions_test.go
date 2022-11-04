package authorization

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/authorization/mgmt/2015-07-01/authorization"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createRoleDefinitionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockAuthorizationRoleDefinitionsClient(ctrl)
	s := services.Services{
		Authorization: services.AuthorizationClient{
			RoleDefinitions: mockClient,
		},
	}

	data := authorization.RoleDefinition{}
	require.Nil(t, faker.FakeObject(&data))

	result := authorization.NewRoleDefinitionListResultPage(authorization.RoleDefinitionListResult{Value: &[]authorization.RoleDefinition{data}}, func(ctx context.Context, result authorization.RoleDefinitionListResult) (authorization.RoleDefinitionListResult, error) {
		return authorization.RoleDefinitionListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), gomock.Any(), "").AnyTimes().Return(result, nil)
	return s
}

func TestAuthorizationRoleDefinitions(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureAuthorizationRoleDefinitionsGenerator{}), createRoleDefinitionsMock, azure_client.TestOptions{})
}
