package frontdoor

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createDoorsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockFrontDoorDoorsClient(ctrl)
	s := services.Services{
		FrontDoor: services.FrontDoorClient{
			Doors: mockClient,
		},
	}

	data := frontdoor.FrontDoor{}
	require.Nil(t, faker.FakeObject(&data))

	result := frontdoor.NewListResultPage(frontdoor.ListResult{Value: &[]frontdoor.FrontDoor{data}}, func(ctx context.Context, result frontdoor.ListResult) (frontdoor.ListResult, error) {
		return frontdoor.ListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestFrontDoorDoors(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureFrontdoorDoorsGenerator{}), createDoorsMock, azure_client.TestOptions{})
}
