package subscriptions

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createTenantsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSubscriptionsTenantsClient(ctrl)
	s := services.Services{
		Subscriptions: services.SubscriptionsClient{
			Tenants: mockClient,
		},
	}

	data := armsubscriptions.TenantIDDescription{}
	require.Nil(t, faker.FakeObject(&data))

	pager := runtime.NewPager(runtime.PagingHandler[armsubscriptions.TenantsClientListResponse]{
		More: func(page armsubscriptions.TenantsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *armsubscriptions.TenantsClientListResponse) (armsubscriptions.TenantsClientListResponse, error) {
			return armsubscriptions.TenantsClientListResponse{
				TenantListResult: armsubscriptions.TenantListResult{
					Value: []*armsubscriptions.TenantIDDescription{&data},
				},
			}, nil
		},
	})

	mockClient.EXPECT().NewListPager(gomock.Any()).AnyTimes().Return(
		pager,
	)
	return s
}

func TestSubscriptionsTenants(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureSubscriptionsTenantsGenerator{}), createTenantsMock, azure_client.TestOptions{})
}
