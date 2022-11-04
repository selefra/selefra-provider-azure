package keyvault

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
)

func createManagedHsmsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockKeyVaultManagedHsmsClient(ctrl)
	s := services.Services{
		KeyVault: services.KeyVaultClient{
			ManagedHsms: mockClient,
		},
	}

	data := keyvault.ManagedHsm{}
	require.Nil(t, faker.FakeObject(&data))

	result := keyvault.NewManagedHsmListResultPage(keyvault.ManagedHsmListResult{Value: &[]keyvault.ManagedHsm{data}}, func(ctx context.Context, result keyvault.ManagedHsmListResult) (keyvault.ManagedHsmListResult, error) {
		return keyvault.ManagedHsmListResult{}, nil
	})

	maxResults := int32(100)
	mockClient.EXPECT().ListBySubscription(gomock.Any(), &maxResults).AnyTimes().Return(result, nil)
	return s
}

func TestKeyVaultManagedHsms(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureKeyvaultManagedHsmsGenerator{}), createManagedHsmsMock, azure_client.TestOptions{})
}
