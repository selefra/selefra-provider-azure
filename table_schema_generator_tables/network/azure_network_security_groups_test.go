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

func createSecurityGroupsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkSecurityGroupsClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			SecurityGroups: mockClient,
		},
	}

	data := network.SecurityGroup{}
	require.Nil(t, faker.FakeObject(&data))

	result := network.NewSecurityGroupListResultPage(network.SecurityGroupListResult{Value: &[]network.SecurityGroup{data}}, func(ctx context.Context, result network.SecurityGroupListResult) (network.SecurityGroupListResult, error) {
		return network.SecurityGroupListResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestNetworkSecurityGroups(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureNetworkSecurityGroupsGenerator{}), createSecurityGroupsMock, azure_client.TestOptions{})
}
