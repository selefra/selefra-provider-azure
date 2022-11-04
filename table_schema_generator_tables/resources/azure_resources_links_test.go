package resources

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2016-09-01/links"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createLinksMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockResourcesLinksClient(ctrl)
	s := services.Services{
		Resources: services.ResourcesClient{
			Links: mockClient,
		},
	}

	data := links.ResourceLink{}
	require.Nil(t, faker.FakeObject(&data))

	result := links.NewResourceLinkResultPage(links.ResourceLinkResult{Value: &[]links.ResourceLink{data}}, func(ctx context.Context, result links.ResourceLinkResult) (links.ResourceLinkResult, error) {
		return links.ResourceLinkResult{}, nil
	})

	mockClient.EXPECT().ListAtSubscription(gomock.Any(), "").AnyTimes().Return(result, nil)
	return s
}

func TestResourcesLinks(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureResourcesLinksGenerator{}), createLinksMock, azure_client.TestOptions{})
}
