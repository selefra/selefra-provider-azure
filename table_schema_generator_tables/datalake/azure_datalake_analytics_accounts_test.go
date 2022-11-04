package datalake

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/datalake/analytics/mgmt/account"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createAnalyticsAccountsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockDataLakeAnalyticsAccountsClient(ctrl)
	s := services.Services{
		DataLake: services.DataLakeClient{
			AnalyticsAccounts: mockClient,
		},
	}

	data := account.DataLakeAnalyticsAccountBasic{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	getData := account.DataLakeAnalyticsAccount{}
	require.Nil(t, faker.FakeObject(&getData))

	result := account.NewDataLakeAnalyticsAccountListResultPage(account.DataLakeAnalyticsAccountListResult{Value: &[]account.DataLakeAnalyticsAccountBasic{data}}, func(ctx context.Context, result account.DataLakeAnalyticsAccountListResult) (account.DataLakeAnalyticsAccountListResult, error) {
		return account.DataLakeAnalyticsAccountListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "", nil, nil, "", "", nil).AnyTimes().Return(result, nil)

	mockClient.EXPECT().Get(gomock.Any(), "test", *data.Name).AnyTimes().Return(getData, nil)
	return s
}

func TestDataLakeAnalyticsAccounts(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureDatalakeAnalyticsAccountsGenerator{}), createAnalyticsAccountsMock, azure_client.TestOptions{})
}
