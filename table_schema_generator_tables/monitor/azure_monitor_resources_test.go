package monitor

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createDiagnosticSettingsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMonitorDiagnosticSettingsClient(ctrl)
	s := services.Services{
		Monitor: services.MonitorClient{
			DiagnosticSettings: mockClient,
		},
	}

	data := insights.DiagnosticSettingsResource{}
	require.Nil(t, faker.FakeObject(&data))

	result := insights.DiagnosticSettingsResourceCollection{Value: &[]insights.DiagnosticSettingsResource{data}}

	mockClient.EXPECT().List(gomock.Any(), "/subscriptions/testSubscription").Return(result, nil)
	mockClient.EXPECT().List(gomock.Any(), "/subscriptions/test/resourceGroups/test/providers/test/test/test").Return(result, nil)
	return s
}

func createResourcesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMonitorResourcesClient(ctrl)
	s := services.Services{
		Monitor: services.MonitorClient{
			Resources:          mockClient,
			DiagnosticSettings: createDiagnosticSettingsMock(t, ctrl).Monitor.DiagnosticSettings,
		},
	}

	data := resources.GenericResourceExpanded{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := resources.NewListResultPage(resources.ListResult{Value: &[]resources.GenericResourceExpanded{data}}, func(ctx context.Context, result resources.ListResult) (resources.ListResult, error) {
		return resources.ListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "", "", nil).AnyTimes().Return(result, nil)
	return s
}

func TestMonitorResources(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureMonitorResourcesGenerator{}), createResourcesMock, azure_client.TestOptions{})
}
