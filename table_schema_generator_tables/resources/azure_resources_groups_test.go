package resources

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createGroupsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockResourcesGroupsClient(ctrl)
	s := services.Services{
		Resources: services.ResourcesClient{
			Groups: mockClient,
		},
	}

	data := resources.Group{}
	require.Nil(t, faker.FakeObject(&data))

	result := resources.NewGroupListResultPage(resources.GroupListResult{Value: &[]resources.Group{data}}, func(ctx context.Context, result resources.GroupListResult) (resources.GroupListResult, error) {
		return resources.GroupListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "", nil).AnyTimes().Return(result, nil)
	return s
}

func TestResourcesGroups(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureResourcesGroupsGenerator{}), createGroupsMock, azure_client.TestOptions{})
}
