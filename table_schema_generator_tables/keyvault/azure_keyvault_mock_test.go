package keyvault

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
)

func createKeysMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockKeyVaultKeysClient(ctrl)
	s := services.Services{
		KeyVault: services.KeyVaultClient{
			Keys: mockClient,
		},
	}

	data := keyvault.KeyItem{}
	require.Nil(t, faker.FakeObject(&data))

	result := keyvault.NewKeyListResultPage(keyvault.KeyListResult{Value: &[]keyvault.KeyItem{data}}, func(ctx context.Context, result keyvault.KeyListResult) (keyvault.KeyListResult, error) {
		return keyvault.KeyListResult{}, nil
	})

	maxResults := int32(25)
	mockClient.EXPECT().GetKeys(gomock.Any(), "test", &maxResults).Return(result, nil)
	return s
}

func createSecretsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockKeyVaultSecretsClient(ctrl)
	s := services.Services{
		KeyVault: services.KeyVaultClient{
			Secrets: mockClient,
		},
	}

	data := keyvault.SecretItem{}
	require.Nil(t, faker.FakeObject(&data))

	result := keyvault.NewSecretListResultPage(keyvault.SecretListResult{Value: &[]keyvault.SecretItem{data}}, func(ctx context.Context, result keyvault.SecretListResult) (keyvault.SecretListResult, error) {
		return keyvault.SecretListResult{}, nil
	})

	maxResults := int32(25)
	mockClient.EXPECT().GetSecrets(gomock.Any(), "test", &maxResults).Return(result, nil)
	return s
}
