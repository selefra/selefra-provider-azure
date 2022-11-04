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

func createAssessmentsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockSecurityAssessmentsClient(ctrl)
	s := services.Services{
		Security: services.SecurityClient{
			Assessments: mockClient,
		},
	}

	data := security.Assessment{}
	require.Nil(t, faker.FakeObject(&data))

	result := security.NewAssessmentListPage(security.AssessmentList{Value: &[]security.Assessment{data}}, func(ctx context.Context, result security.AssessmentList) (security.AssessmentList, error) {
		return security.AssessmentList{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "/subscriptions/testSubscription").AnyTimes().Return(result, nil)
	return s
}

func TestSecurityAssessments(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureSecurityAssessmentsGenerator{}), createAssessmentsMock, azure_client.TestOptions{})
}
