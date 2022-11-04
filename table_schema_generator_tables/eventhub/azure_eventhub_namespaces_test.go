package eventhub

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/eventhub/mgmt/2018-01-01-preview/eventhub"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createNetworkRuleSetsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockEventHubNetworkRuleSetsClient(ctrl)
	s := services.Services{
		EventHub: services.EventHubClient{
			NetworkRuleSets: mockClient,
		},
	}

	data := eventhub.NetworkRuleSet{}
	require.Nil(t, faker.FakeObject(&data))

	result := data

	mockClient.EXPECT().GetNetworkRuleSet(gomock.Any(), "test", "test").Return(result, nil)
	return s
}

func createNamespacesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockEventHubNamespacesClient(ctrl)
	s := services.Services{
		EventHub: services.EventHubClient{
			Namespaces:      mockClient,
			NetworkRuleSets: createNetworkRuleSetsMock(t, ctrl).EventHub.NetworkRuleSets,
		},
	}

	data := eventhub.EHNamespace{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := eventhub.NewEHNamespaceListResultPage(eventhub.EHNamespaceListResult{Value: &[]eventhub.EHNamespace{data}}, func(ctx context.Context, result eventhub.EHNamespaceListResult) (eventhub.EHNamespaceListResult, error) {
		return eventhub.EHNamespaceListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestEventHubNamespaces(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureEventhubNamespacesGenerator{}), createNamespacesMock, azure_client.TestOptions{})
}
