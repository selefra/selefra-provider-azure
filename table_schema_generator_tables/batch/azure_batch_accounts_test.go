package batch

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/batch/mgmt/2021-06-01/batch"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createAccountsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockBatchAccountsClient(ctrl)
	s := services.Services{
		Batch: services.BatchClient{
			Accounts: mockClient,
		},
	}

	data := batch.Account{}
	require.Nil(t, faker.FakeObject(&data))

	result := batch.NewAccountListResultPage(batch.AccountListResult{Value: &[]batch.Account{data}}, func(ctx context.Context, result batch.AccountListResult) (batch.AccountListResult, error) {
		return batch.AccountListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestBatchAccounts(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureBatchAccountsGenerator{}), createAccountsMock, azure_client.TestOptions{})
}
