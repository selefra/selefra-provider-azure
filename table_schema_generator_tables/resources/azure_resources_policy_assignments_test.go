package resources

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2020-03-01-preview/policy"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createPolicyAssignmentsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockResourcesPolicyAssignmentsClient(ctrl)
	s := services.Services{
		Resources: services.ResourcesClient{
			PolicyAssignments: mockClient,
		},
	}

	data := policy.Assignment{}
	require.Nil(t, faker.FakeObject(&data))

	result := policy.NewAssignmentListResultPage(policy.AssignmentListResult{Value: &[]policy.Assignment{data}}, func(ctx context.Context, result policy.AssignmentListResult) (policy.AssignmentListResult, error) {
		return policy.AssignmentListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), gomock.Any(), "", nil).AnyTimes().Return(result, nil)
	return s
}

func TestResourcesPolicyAssignments(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureResourcesPolicyAssignmentsGenerator{}), createPolicyAssignmentsMock, azure_client.TestOptions{})
}
