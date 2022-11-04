package security

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createPricingsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSecurityPricingsClient(ctrl)
	s := services.Services{
		Security: services.SecurityClient{
			Pricings: mockClient,
		},
	}

	data := security.Pricing{}
	require.Nil(t, faker.FakeObject(&data))

	result := security.PricingList{Value: &[]security.Pricing{data}}

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestSecurityPricings(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureSecurityPricingsGenerator{}), createPricingsMock, azure_client.TestOptions{})
}
