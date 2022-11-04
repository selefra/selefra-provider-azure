package security

import (
	"context"
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

func createSettingsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSecuritySettingsClient(ctrl)
	s := services.Services{
		Security: services.SecurityClient{
			Settings: mockClient,
		},
	}

	data := security.Setting{}
	require.Nil(t, faker.FakeObject(&data))

	result := security.NewSettingsListPage(security.SettingsList{Value: &[]security.BasicSetting{data}}, func(ctx context.Context, result security.SettingsList) (security.SettingsList, error) {
		return security.SettingsList{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestSecuritySettings(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureSecuritySettingsGenerator{}), createSettingsMock, azure_client.TestOptions{})
}
