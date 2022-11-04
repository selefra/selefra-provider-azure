package container

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-03-01/containerservice"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createManagedClustersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockContainerManagedClustersClient(ctrl)
	s := services.Services{
		Container: services.ContainerClient{
			ManagedClusters: mockClient,
		},
	}

	data := containerservice.ManagedCluster{}
	require.Nil(t, faker.FakeObject(&data))

	result := containerservice.NewManagedClusterListResultPage(containerservice.ManagedClusterListResult{Value: &[]containerservice.ManagedCluster{data}}, func(ctx context.Context, result containerservice.ManagedClusterListResult) (containerservice.ManagedClusterListResult, error) {
		return containerservice.ManagedClusterListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestContainerManagedClusters(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureContainerManagedClustersGenerator{}), createManagedClustersMock, azure_client.TestOptions{})
}
