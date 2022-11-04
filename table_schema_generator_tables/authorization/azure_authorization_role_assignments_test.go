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

func createRoleAssignmentsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockAuthorizationRoleAssignmentsClient(ctrl)
	s := services.Services{
		Authorization: services.AuthorizationClient{
			RoleAssignments: mockClient,
		},
	}

	data := authorization.RoleAssignment{}
	require.Nil(t, faker.FakeObject(&data))

	result := authorization.NewRoleAssignmentListResultPage(authorization.RoleAssignmentListResult{Value: &[]authorization.RoleAssignment{data}}, func(ctx context.Context, result authorization.RoleAssignmentListResult) (authorization.RoleAssignmentListResult, error) {
		return authorization.RoleAssignmentListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "").AnyTimes().Return(result, nil)
	return s
}

func TestAuthorizationRoleAssignments(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureAuthorizationRoleAssignmentsGenerator{}), createRoleAssignmentsMock, azure_client.TestOptions{})
}
