package datalake

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/store/mgmt/account"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createStoreAccountsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockDataLakeStoreAccountsClient(ctrl)
	s := services.Services{
		DataLake: services.DataLakeClient{
			StoreAccounts: mockClient,
		},
	}

	data := account.DataLakeStoreAccountBasic{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	getData := account.DataLakeStoreAccount{}
	require.Nil(t, faker.FakeObject(&getData))

	result := account.NewDataLakeStoreAccountListResultPage(account.DataLakeStoreAccountListResult{Value: &[]account.DataLakeStoreAccountBasic{data}}, func(ctx context.Context, result account.DataLakeStoreAccountListResult) (account.DataLakeStoreAccountListResult, error) {
		return account.DataLakeStoreAccountListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "", nil, nil, "", "", nil).AnyTimes().Return(result, nil)

	mockClient.EXPECT().Get(gomock.Any(), "test", *data.Name).AnyTimes().Return(getData, nil)
	return s
}

func TestDataLakeStoreAccounts(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureDatalakeStoreAccountsGenerator{}), createStoreAccountsMock, azure_client.TestOptions{})
}
