package monitor

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createLogProfilesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMonitorLogProfilesClient(ctrl)
	s := services.Services{
		Monitor: services.MonitorClient{
			LogProfiles: mockClient,
		},
	}

	data := insights.LogProfileResource{}
	require.Nil(t, faker.FakeObject(&data))

	result := insights.LogProfileCollection{Value: &[]insights.LogProfileResource{data}}

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestMonitorLogProfiles(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureMonitorLogProfilesGenerator{}), createLogProfilesMock, azure_client.TestOptions{})
}
