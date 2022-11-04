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

func createExpressRouteCircuitsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkExpressRouteCircuitsClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			ExpressRouteCircuits: mockClient,
		},
	}

	data := network.ExpressRouteCircuit{}
	require.Nil(t, faker.FakeObject(&data))

	result := network.NewExpressRouteCircuitListResultPage(network.ExpressRouteCircuitListResult{Value: &[]network.ExpressRouteCircuit{data}}, func(ctx context.Context, result network.ExpressRouteCircuitListResult) (network.ExpressRouteCircuitListResult, error) {
		return network.ExpressRouteCircuitListResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestNetworkExpressRouteCircuits(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureNetworkExpressRouteCircuitsGenerator{}), createExpressRouteCircuitsMock, azure_client.TestOptions{})
}
