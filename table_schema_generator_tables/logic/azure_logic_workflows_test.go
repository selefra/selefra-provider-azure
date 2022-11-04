package logic

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/monitor/mgmt/insights"
	"github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2019-05-01/logic"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createDiagnosticSettingsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockLogicDiagnosticSettingsClient(ctrl)
	s := services.Services{
		Logic: services.LogicClient{
			DiagnosticSettings: mockClient,
		},
	}

	data := insights.DiagnosticSettingsResource{}
	require.Nil(t, faker.FakeObject(&data))

	result := insights.DiagnosticSettingsResourceCollection{Value: &[]insights.DiagnosticSettingsResource{data}}

	mockClient.EXPECT().List(gomock.Any(), "/subscriptions/test/resourceGroups/test/providers/test/test/test").Return(result, nil)
	return s
}

func createWorkflowsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockLogicWorkflowsClient(ctrl)
	s := services.Services{
		Logic: services.LogicClient{
			Workflows:          mockClient,
			DiagnosticSettings: createDiagnosticSettingsMock(t, ctrl).Logic.DiagnosticSettings,
		},
	}

	data := logic.Workflow{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := logic.NewWorkflowListResultPage(logic.WorkflowListResult{Value: &[]logic.Workflow{data}}, func(ctx context.Context, result logic.WorkflowListResult) (logic.WorkflowListResult, error) {
		return logic.WorkflowListResult{}, nil
	})

	var top int32 = 100
	mockClient.EXPECT().ListBySubscription(gomock.Any(), &top, "").AnyTimes().Return(result, nil)
	return s
}

func TestLogicWorkflows(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureLogicWorkflowsGenerator{}), createWorkflowsMock, azure_client.TestOptions{})
}
