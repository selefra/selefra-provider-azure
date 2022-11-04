package iothub

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/iothub/mgmt/2021-07-02/devices"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createDevicesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockIotHubDevicesClient(ctrl)
	s := services.Services{
		IotHub: services.IotHubClient{
			Devices: mockClient,
		},
	}

	data := devices.IotHubDescription{}
	require.Nil(t, faker.FakeObject(&data))

	result := devices.NewIotHubDescriptionListResultPage(devices.IotHubDescriptionListResult{Value: &[]devices.IotHubDescription{data}}, func(ctx context.Context, result devices.IotHubDescriptionListResult) (devices.IotHubDescriptionListResult, error) {
		return devices.IotHubDescriptionListResult{}, nil
	})

	mockClient.EXPECT().ListBySubscription(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestIotHubDevices(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureIothubDevicesGenerator{}), createDevicesMock, azure_client.TestOptions{})
}
