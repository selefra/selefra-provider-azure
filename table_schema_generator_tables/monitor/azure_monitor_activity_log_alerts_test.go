package monitor

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-11-01-preview/insights"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createActivityLogAlertsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMonitorActivityLogAlertsClient(ctrl)
	s := services.Services{
		Monitor: services.MonitorClient{
			ActivityLogAlerts: mockClient,
		},
	}

	data := insights.ActivityLogAlertResource{}
	require.Nil(t, faker.FakeObject(&data))

	result := insights.ActivityLogAlertList{Value: &[]insights.ActivityLogAlertResource{data}}

	mockClient.EXPECT().ListBySubscriptionID(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestMonitorActivityLogAlerts(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureMonitorActivityLogAlertsGenerator{}), createActivityLogAlertsMock, azure_client.TestOptions{})
}
