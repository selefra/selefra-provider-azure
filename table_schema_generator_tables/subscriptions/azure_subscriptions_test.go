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

func createSubscriptionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSubscriptionsSubscriptionsClient(ctrl)
	s := services.Services{
		Subscriptions: services.SubscriptionsClient{
			Subscriptions: mockClient,
		},
	}

	data := armsubscriptions.Subscription{}
	require.Nil(t, faker.FakeObject(&data))

	pager := runtime.NewPager(runtime.PagingHandler[armsubscriptions.ClientListResponse]{
		More: func(page armsubscriptions.ClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *armsubscriptions.ClientListResponse) (armsubscriptions.ClientListResponse, error) {
			return armsubscriptions.ClientListResponse{
				SubscriptionListResult: armsubscriptions.SubscriptionListResult{
					Value: []*armsubscriptions.Subscription{&data},
				},
			}, nil
		},
	})

	mockClient.EXPECT().NewListPager(gomock.Any()).AnyTimes().Return(
		pager,
	)
	return s
}

func TestSubscriptionsSubscriptions(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureSubscriptionsGenerator{}), createSubscriptionsMock, azure_client.TestOptions{})
}
