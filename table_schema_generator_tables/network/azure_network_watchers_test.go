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

func createFlowLogsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkFlowLogsClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			FlowLogs: mockClient,
		},
	}

	data := network.FlowLog{}
	require.Nil(t, faker.FakeObject(&data))

	result := network.NewFlowLogListResultPage(network.FlowLogListResult{Value: &[]network.FlowLog{data}}, func(ctx context.Context, result network.FlowLogListResult) (network.FlowLogListResult, error) {
		return network.FlowLogListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "test", "test").Return(result, nil)
	return s
}

func createWatchersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockNetworkWatchersClient(ctrl)
	s := services.Services{
		Network: services.NetworkClient{
			Watchers: mockClient,
			FlowLogs: createFlowLogsMock(t, ctrl).Network.FlowLogs,
		},
	}

	data := network.Watcher{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := network.WatcherListResult{Value: &[]network.Watcher{data}}

	mockClient.EXPECT().ListAll(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestNetworkWatchers(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureNetworkWatchersGenerator{}), createWatchersMock, azure_client.TestOptions{})
}
