package network

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createInterfacesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkInterfacesClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			Interfaces: mockClient,
		},
	}

	data := network.Interface{}
	require.Nil(t, faker.FakeObject(&data))

	result := network.NewInterfaceListResultPage(network.InterfaceListResult{Value: &[]network.Interface{data}}, func(ctx context.Context, result network.InterfaceListResult) (network.InterfaceListResult, error) {
		return network.InterfaceListResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestNetworkInterfaces(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureNetworkInterfacesGenerator{}), createInterfacesMock, azure_client.TestOptions{})
}
