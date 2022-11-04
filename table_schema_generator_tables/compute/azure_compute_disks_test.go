package compute

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createDisksMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockComputeDisksClient(ctrl)
	s := services.Services{
		Compute: services.ComputeClient{
			Disks: mockClient,
		},
	}

	data := compute.Disk{}
	require.Nil(t, faker.FakeObject(&data))

	result := compute.NewDiskListPage(compute.DiskList{Value: &[]compute.Disk{data}}, func(ctx context.Context, result compute.DiskList) (compute.DiskList, error) {
		return compute.DiskList{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestComputeDisks(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureComputeDisksGenerator{}), createDisksMock, azure_client.TestOptions{})
}
