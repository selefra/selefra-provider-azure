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

func createVirtualNetworkGatewaysMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkVirtualNetworkGatewaysClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			VirtualNetworkGateways: mockClient,
		},
	}

	data := network.VirtualNetworkGateway{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := network.NewVirtualNetworkGatewayListResultPage(network.VirtualNetworkGatewayListResult{Value: &[]network.VirtualNetworkGateway{data}}, func(ctx context.Context, result network.VirtualNetworkGatewayListResult) (network.VirtualNetworkGatewayListResult, error) {
		return network.VirtualNetworkGatewayListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "test").Return(result, nil)
	return s
}
func createVirtualNetworkGatewayConnectionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkVirtualNetworkGatewayConnectionsClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			VirtualNetworkGatewayConnections: mockClient,
		},
	}

	data := network.VirtualNetworkGatewayConnectionListEntity{}
	require.Nil(t, faker.FakeObject(&data))

	result := network.NewVirtualNetworkGatewayListConnectionsResultPage(network.VirtualNetworkGatewayListConnectionsResult{Value: &[]network.VirtualNetworkGatewayConnectionListEntity{data}}, func(ctx context.Context, result network.VirtualNetworkGatewayListConnectionsResult) (network.VirtualNetworkGatewayListConnectionsResult, error) {
		return network.VirtualNetworkGatewayListConnectionsResult{}, nil
	})

	mockClient.EXPECT().ListConnections(gomock.Any(), "test", "test").Return(result, nil)
	return s
}

func createVirtualNetworksMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkVirtualNetworksClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			VirtualNetworks:                  mockClient,
			VirtualNetworkGateways:           createVirtualNetworkGatewaysMock(t, ctrl).Network.VirtualNetworkGateways,
			VirtualNetworkGatewayConnections: createVirtualNetworkGatewayConnectionsMock(t, ctrl).Network.VirtualNetworkGatewayConnections,
		},
	}

	data := network.VirtualNetwork{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := network.NewVirtualNetworkListResultPage(network.VirtualNetworkListResult{Value: &[]network.VirtualNetwork{data}}, func(ctx context.Context, result network.VirtualNetworkListResult) (network.VirtualNetworkListResult, error) {
		return network.VirtualNetworkListResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestNetworkVirtualNetworks(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureNetworkVirtualNetworksGenerator{}), createVirtualNetworksMock, azure_client.TestOptions{})
}
