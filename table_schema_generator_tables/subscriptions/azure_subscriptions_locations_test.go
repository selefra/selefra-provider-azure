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

func createLocationsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSubscriptionsLocationsClient(ctrl)
	s := services.Services{
		Subscriptions: services.SubscriptionsClient{
			Locations: mockClient,
		},
	}

	data := armsubscriptions.Location{}
	require.Nil(t, faker.FakeObject(&data))

	pager := runtime.NewPager(runtime.PagingHandler[armsubscriptions.ClientListLocationsResponse]{
		More: func(page armsubscriptions.ClientListLocationsResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *armsubscriptions.ClientListLocationsResponse) (armsubscriptions.ClientListLocationsResponse, error) {
			return armsubscriptions.ClientListLocationsResponse{
				LocationListResult: armsubscriptions.LocationListResult{
					Value: []*armsubscriptions.Location{&data},
				},
			}, nil
		},
	})

	mockClient.EXPECT().NewListLocationsPager(gomock.Any(), gomock.Any()).AnyTimes().Return(
		pager,
	)
	return s
}

func TestSubscriptionsLocations(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureSubscriptionsLocationsGenerator{}), createLocationsMock, azure_client.TestOptions{})
}
