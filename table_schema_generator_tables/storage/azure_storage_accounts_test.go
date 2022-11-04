package storage

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
	"github.com/tombuildsstuff/giovanni/storage/2020-08-04/blob/accounts"
	"github.com/tombuildsstuff/giovanni/storage/2020-08-04/queue/queues"
)

func createContainersMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockStorageContainersClient(ctrl)
	s := services.Services{
		Storage: services.StorageClient{
			Containers: mockClient,
		},
	}

	data := storage.ListContainerItem{}
	require.Nil(t, faker.FakeObject(&data))

	result := storage.NewListContainerItemsPage(storage.ListContainerItems{Value: &[]storage.ListContainerItem{data}}, func(ctx context.Context, result storage.ListContainerItems) (storage.ListContainerItems, error) {
		return storage.ListContainerItems{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "test", "test", "", "", gomock.Any()).Return(result, nil)
	return s
}
func createBlobServicesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockStorageBlobServicesClient(ctrl)
	s := services.Services{
		Storage: services.StorageClient{
			BlobServices: mockClient,
		},
	}

	data := storage.BlobServiceProperties{}
	require.Nil(t, faker.FakeObject(&data))

	result := storage.BlobServiceItems{Value: &[]storage.BlobServiceProperties{data}}

	mockClient.EXPECT().List(gomock.Any(), "test", "test").Return(result, nil)
	return s
}

func createAccountsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockStorageAccountsClient(ctrl)
	s := services.Services{
		Storage: services.StorageClient{
			Accounts:     mockClient,
			Containers:   createContainersMock(t, ctrl).Storage.Containers,
			BlobServices: createBlobServicesMock(t, ctrl).Storage.BlobServices,
		},
	}

	data := storage.Account{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := storage.NewAccountListResultPage(storage.AccountListResult{Value: &[]storage.Account{data}}, func(ctx context.Context, result storage.AccountListResult) (storage.AccountListResult, error) {
		return storage.AccountListResult{}, nil
	})

	result.Values()[0].Sku.Tier = storage.Standard
	result.Values()[0].Kind = storage.StorageV2
	blobProperties := accounts.StorageServiceProperties{}
	require.Nil(t, faker.FakeObject(&blobProperties))
	blobResult := accounts.GetServicePropertiesResult{StorageServiceProperties: &blobProperties}
	mockClient.EXPECT().GetBlobServiceProperties(gomock.Any(), "test", "test").AnyTimes().Return(blobResult, nil)
	queueProperties := queues.StorageServiceProperties{}
	require.Nil(t, faker.FakeObject(&queueProperties))
	queueResult := queues.StorageServicePropertiesResponse{StorageServiceProperties: queueProperties}
	mockClient.EXPECT().GetQueueServiceProperties(gomock.Any(), "test", "test").AnyTimes().Return(queueResult, nil)
	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestStorageAccounts(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureStorageAccountsGenerator{}), createAccountsMock, azure_client.TestOptions{})
}
