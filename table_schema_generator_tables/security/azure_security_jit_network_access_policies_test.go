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

func createJitNetworkAccessPoliciesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSecurityJitNetworkAccessPoliciesClient(ctrl)
	s := services.Services{
		Security: services.SecurityClient{
			JitNetworkAccessPolicies: mockClient,
		},
	}

	data := security.JitNetworkAccessPolicy{}
	require.Nil(t, faker.FakeObject(&data))

	result := security.NewJitNetworkAccessPoliciesListPage(security.JitNetworkAccessPoliciesList{Value: &[]security.JitNetworkAccessPolicy{data}}, func(ctx context.Context, result security.JitNetworkAccessPoliciesList) (security.JitNetworkAccessPoliciesList, error) {
		return security.JitNetworkAccessPoliciesList{}, nil
	})

	mockClient.EXPECT().List(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestSecurityJitNetworkAccessPolicies(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureSecurityJitNetworkAccessPoliciesGenerator{}), createJitNetworkAccessPoliciesMock, azure_client.TestOptions{})
}
