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

func createPublicIPAddressesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkPublicIPAddressesClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			PublicIPAddresses: mockClient,
		},
	}

	data := network.PublicIPAddress{}
	require.Nil(t, faker.FakeObject(&data))

	result := network.NewPublicIPAddressListResultPage(network.PublicIPAddressListResult{Value: &[]network.PublicIPAddress{data}}, func(ctx context.Context, result network.PublicIPAddressListResult) (network.PublicIPAddressListResult, error) {
		return network.PublicIPAddressListResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestNetworkPublicIPAddresses(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureNetworkPublicIpAddressesGenerator{}), createPublicIPAddressesMock, azure_client.TestOptions{})
}
