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

func createInstanceViewsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockComputeInstanceViewsClient(ctrl)
	s := services.Services{
		Compute: services.ComputeClient{
			InstanceViews: mockClient,
		},
	}

	data := compute.VirtualMachineInstanceView{}
	require.Nil(t, faker.FakeObject(&data))

	result := data

	mockClient.EXPECT().InstanceView(gomock.Any(), "test", "test").Return(result, nil)
	return s
}
func createVirtualMachineExtensionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockComputeVirtualMachineExtensionsClient(ctrl)
	s := services.Services{
		Compute: services.ComputeClient{
			VirtualMachineExtensions: mockClient,
		},
	}

	data := compute.VirtualMachineExtension{}
	require.Nil(t, faker.FakeObject(&data))

	result := compute.VirtualMachineExtensionsListResult{Value: &[]compute.VirtualMachineExtension{data}}

	mockClient.EXPECT().List(gomock.Any(), "test", "test", "").Return(result, nil)
	return s
}

func createVirtualMachinesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockComputeVirtualMachinesClient(ctrl)
	s := services.Services{
		Compute: services.ComputeClient{
			VirtualMachines:          mockClient,
			InstanceViews:            createInstanceViewsMock(t, ctrl).Compute.InstanceViews,
			VirtualMachineExtensions: createVirtualMachineExtensionsMock(t, ctrl).Compute.VirtualMachineExtensions,
		},
	}

	data := compute.VirtualMachine{}
	require.Nil(t, faker.FakeObject(&data))

	name := "test"
	data.Name = &name

	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := compute.NewVirtualMachineListResultPage(compute.VirtualMachineListResult{Value: &[]compute.VirtualMachine{data}}, func(ctx context.Context, result compute.VirtualMachineListResult) (compute.VirtualMachineListResult, error) {
		return compute.VirtualMachineListResult{}, nil
	})

	mockClient.EXPECT().ListAll(gomock.Any(), "false").AnyTimes().Return(result, nil)
	return s
}

func TestComputeVirtualMachines(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureComputeVirtualMachinesGenerator{}), createVirtualMachinesMock, azure_client.TestOptions{})
}
