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

func createVirtualMachineScaleSetsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockComputeVirtualMachineScaleSetsClient(ctrl)
	s := services.Services{
		Compute: services.ComputeClient{
			VirtualMachineScaleSets: mockClient,
		},
	}

	data := compute.VirtualMachineScaleSet{}
	require.Nil(t, faker.FakeObject(&data))

	result := compute.NewVirtualMachineScaleSetListWithLinkResultPage(compute.VirtualMachineScaleSetListWithLinkResult{Value: &[]compute.VirtualMachineScaleSet{data}}, func(ctx context.Context, result compute.VirtualMachineScaleSetListWithLinkResult) (compute.VirtualMachineScaleSetListWithLinkResult, error) {
		return compute.VirtualMachineScaleSetListWithLinkResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestComputeVirtualMachineScaleSets(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureComputeVirtualMachineScaleSetsGenerator{}), createVirtualMachineScaleSetsMock, azure_client.TestOptions{})
}
