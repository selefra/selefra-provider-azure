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

func createRouteTablesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkRouteTablesClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			RouteTables: mockClient,
		},
	}

	data := network.RouteTable{}
	require.Nil(t, faker.FakeObject(&data))

	result := network.NewRouteTableListResultPage(network.RouteTableListResult{Value: &[]network.RouteTable{data}}, func(ctx context.Context, result network.RouteTableListResult) (network.RouteTableListResult, error) {
		return network.RouteTableListResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestNetworkRouteTables(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureNetworkRouteTablesGenerator{}), createRouteTablesMock, azure_client.TestOptions{})
}
