package web

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createSiteAuthSettingsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockWebSiteAuthSettingsClient(ctrl)
	s := services.Services{
		Web: services.WebClient{
			SiteAuthSettings: mockClient,
		},
	}

	data := web.SiteAuthSettings{}
	require.Nil(t, faker.FakeObject(&data))

	result := data

	mockClient.EXPECT().GetAuthSettings(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createVnetConnectionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockWebVnetConnectionsClient(ctrl)
	s := services.Services{
		Web: services.WebClient{
			VnetConnections: mockClient,
		},
	}

	data := web.VnetInfo{}
	require.Nil(t, faker.FakeObject(&data))

	result := data

	mockClient.EXPECT().GetVnetConnection(gomock.Any(), "test", "test", "test").Return(result, nil)
	return s
}
func createPublishingProfilesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockWebPublishingProfilesClient(ctrl)
	s := services.Services{
		Web: services.WebClient{
			PublishingProfiles: mockClient,
		},
	}

	data := services.PublishingProfiles{}
	require.Nil(t, faker.FakeObject(&data))

	result := data

	mockClient.EXPECT().ListPublishingProfiles(gomock.Any(), "test", "test").Return(result, nil)
	return s
}

func createAppsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockWebAppsClient(ctrl)
	s := services.Services{
		Web: services.WebClient{
			Apps:               mockClient,
			SiteAuthSettings:   createSiteAuthSettingsMock(t, ctrl).Web.SiteAuthSettings,
			VnetConnections:    createVnetConnectionsMock(t, ctrl).Web.VnetConnections,
			PublishingProfiles: createPublishingProfilesMock(t, ctrl).Web.PublishingProfiles,
		},
	}

	data := web.Site{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := web.NewAppCollectionPage(web.AppCollection{Value: &[]web.Site{data}}, func(ctx context.Context, result web.AppCollection) (web.AppCollection, error) {
		return web.AppCollection{}, nil
	})

	vnetName := "test"
	result.Values()[0].SiteConfig.VnetName = &vnetName
	resourceGroup := "test"
	result.Values()[0].ResourceGroup = &resourceGroup
	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestWebApps(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureWebAppsGenerator{}), createAppsMock, azure_client.TestOptions{})
}
