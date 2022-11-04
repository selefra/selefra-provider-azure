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

func createAutoProvisioningSettingsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSecurityAutoProvisioningSettingsClient(ctrl)
	s := services.Services{
		Security: services.SecurityClient{
			AutoProvisioningSettings: mockClient,
		},
	}

	data := security.AutoProvisioningSetting{}
	require.Nil(t, faker.FakeObject(&data))

	result := security.NewAutoProvisioningSettingListPage(security.AutoProvisioningSettingList{Value: &[]security.AutoProvisioningSetting{data}}, func(ctx context.Context, result security.AutoProvisioningSettingList) (security.AutoProvisioningSettingList, error) {
		return security.AutoProvisioningSettingList{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestSecurityAutoProvisioningSettings(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureSecurityAutoProvisioningSettingsGenerator{}), createAutoProvisioningSettingsMock, azure_client.TestOptions{})
}
