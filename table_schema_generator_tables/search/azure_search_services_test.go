package search

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/search/mgmt/2020-08-01/search"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createServicesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSearchServicesClient(ctrl)
	s := services.Services{
		Search: services.SearchClient{
			AzureServices: mockClient,
		},
	}

	data := search.Service{}
	require.Nil(t, faker.FakeObject(&data))

	result := search.NewServiceListResultPage(search.ServiceListResult{Value: &[]search.Service{data}}, func(ctx context.Context, result search.ServiceListResult) (search.ServiceListResult, error) {
		return search.ServiceListResult{}, nil
	})

	mockClient.EXPECT().ListBySubscription(gomock.Any(), nil).AnyTimes().Return(result, nil)
	return s
}

func TestSearchServices(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureSearchServicesGenerator{}), createServicesMock, azure_client.TestOptions{})
}
