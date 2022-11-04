package network

import (
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

func createExpressRouteGatewaysMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkExpressRouteGatewaysClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			ExpressRouteGateways: mockClient,
		},
	}

	data := network.ExpressRouteGateway{}
	require.Nil(t, faker.FakeObject(&data))

	result := network.ExpressRouteGatewayList{Value: &[]network.ExpressRouteGateway{data}}

	mockClient.EXPECT().ListBySubscription(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestNetworkExpressRouteGateways(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureNetworkExpressRouteGatewaysGenerator{}), createExpressRouteGatewaysMock, azure_client.TestOptions{})
}
