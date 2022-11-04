package servicebus

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createTopicsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockServicebusTopicsClient(ctrl)
	s := services.Services{
		Servicebus: services.ServicebusClient{
			Topics: mockClient,
		},
	}

	data := servicebus.SBTopic{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := servicebus.NewSBTopicListResultPage(servicebus.SBTopicListResult{Value: &[]servicebus.SBTopic{data}}, func(ctx context.Context, result servicebus.SBTopicListResult) (servicebus.SBTopicListResult, error) {
		return servicebus.SBTopicListResult{}, nil
	})

	mockClient.EXPECT().ListByNamespace(gomock.Any(), "test", "test", nil, nil).Return(result, nil)
	return s
}
func createAuthorizationRulesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockServicebusAuthorizationRulesClient(ctrl)
	s := services.Services{
		Servicebus: services.ServicebusClient{
			AuthorizationRules: mockClient,
		},
	}

	data := servicebus.SBAuthorizationRule{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := servicebus.NewSBAuthorizationRuleListResultPage(servicebus.SBAuthorizationRuleListResult{Value: &[]servicebus.SBAuthorizationRule{data}}, func(ctx context.Context, result servicebus.SBAuthorizationRuleListResult) (servicebus.SBAuthorizationRuleListResult, error) {
		return servicebus.SBAuthorizationRuleListResult{}, nil
	})

	mockClient.EXPECT().ListAuthorizationRules(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}
func createAccessKeysMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockServicebusAccessKeysClient(ctrl)
	s := services.Services{
		Servicebus: services.ServicebusClient{
			AccessKeys: mockClient,
		},
	}

	data := servicebus.AccessKeys{}
	require.Nil(t, faker.FakeObject(&data))

	result := data

	mockClient.EXPECT().ListKeys(gomock.Any(), "test", "test", "test", "test").Return(result, nil)
	return s
}

func createNamespacesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockServicebusNamespacesClient(ctrl)
	s := services.Services{
		Servicebus: services.ServicebusClient{
			Namespaces:         mockClient,
			Topics:             createTopicsMock(t, ctrl).Servicebus.Topics,
			AuthorizationRules: createAuthorizationRulesMock(t, ctrl).Servicebus.AuthorizationRules,
			AccessKeys:         createAccessKeysMock(t, ctrl).Servicebus.AccessKeys,
		},
	}

	data := servicebus.SBNamespace{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := servicebus.NewSBNamespaceListResultPage(servicebus.SBNamespaceListResult{Value: &[]servicebus.SBNamespace{data}}, func(ctx context.Context, result servicebus.SBNamespaceListResult) (servicebus.SBNamespaceListResult, error) {
		return servicebus.SBNamespaceListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestServicebusNamespaces(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureServicebusNamespacesGenerator{}), createNamespacesMock, azure_client.TestOptions{})
}
