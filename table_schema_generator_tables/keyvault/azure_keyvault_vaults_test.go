package keyvault

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/keyvault/mgmt/keyvault"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func createVaultsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockKeyVaultVaultsClient(ctrl)
	s := services.Services{
		KeyVault: services.KeyVaultClient{
			Vaults:  mockClient,
			Keys:    createKeysMock(t, ctrl).KeyVault.Keys,
			Secrets: createSecretsMock(t, ctrl).KeyVault.Secrets,
		},
	}

	data := keyvault.Vault{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := keyvault.NewVaultListResultPage(keyvault.VaultListResult{Value: &[]keyvault.Vault{data}}, func(ctx context.Context, result keyvault.VaultListResult) (keyvault.VaultListResult, error) {
		return keyvault.VaultListResult{}, nil
	})

	maxResults := int32(1000)
	vaultURI := "test"
	data.Properties.VaultURI = &vaultURI
	mockClient.EXPECT().ListBySubscription(gomock.Any(), &maxResults).AnyTimes().Return(result, nil)
	return s
}

func TestKeyVaultVaults(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureKeyvaultVaultsGenerator{}), createVaultsMock, azure_client.TestOptions{})
}
