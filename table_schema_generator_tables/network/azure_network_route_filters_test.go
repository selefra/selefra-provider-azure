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

func createRouteFiltersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkRouteFiltersClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			RouteFilters: mockClient,
		},
	}

	data := network.RouteFilter{}
	require.Nil(t, faker.FakeObject(&data))

	result := network.NewRouteFilterListResultPage(network.RouteFilterListResult{Value: &[]network.RouteFilter{data}}, func(ctx context.Context, result network.RouteFilterListResult) (network.RouteFilterListResult, error) {
		return network.RouteFilterListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestNetworkRouteFilters(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureNetworkRouteFiltersGenerator{}), createRouteFiltersMock, azure_client.TestOptions{})
}
