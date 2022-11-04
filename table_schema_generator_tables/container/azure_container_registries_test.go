package container

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/containerregistry/mgmt/2019-05-01/containerregistry"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createReplicationsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockContainerReplicationsClient(ctrl)
	s := services.Services{
		Container: services.ContainerClient{
			Replications: mockClient,
		},
	}

	data := containerregistry.Replication{}
	require.Nil(t, faker.FakeObject(&data))

	result := containerregistry.NewReplicationListResultPage(containerregistry.ReplicationListResult{Value: &[]containerregistry.Replication{data}}, func(ctx context.Context, result containerregistry.ReplicationListResult) (containerregistry.ReplicationListResult, error) {
		return containerregistry.ReplicationListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "test", "test").Return(result, nil)
	return s
}

func createRegistriesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockContainerRegistriesClient(ctrl)
	s := services.Services{
		Container: services.ContainerClient{
			Registries:   mockClient,
			Replications: createReplicationsMock(t, ctrl).Container.Replications,
		},
	}

	data := containerregistry.Registry{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := containerregistry.NewRegistryListResultPage(containerregistry.RegistryListResult{Value: &[]containerregistry.Registry{data}}, func(ctx context.Context, result containerregistry.RegistryListResult) (containerregistry.RegistryListResult, error) {
		return containerregistry.RegistryListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestContainerRegistries(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureContainerRegistriesGenerator{}), createRegistriesMock, azure_client.TestOptions{})
}
