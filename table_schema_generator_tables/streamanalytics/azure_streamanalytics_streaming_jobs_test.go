package streamanalytics

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/streamanalytics/mgmt/2020-03-01/streamanalytics"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createStreamingJobsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockStreamAnalyticsStreamingJobsClient(ctrl)
	s := services.Services{
		StreamAnalytics: services.StreamAnalyticsClient{
			StreamingJobs: mockClient,
		},
	}

	data := streamanalytics.StreamingJob{}
	require.Nil(t, faker.FakeObject(&data))

	result := streamanalytics.NewStreamingJobListResultPage(streamanalytics.StreamingJobListResult{Value: &[]streamanalytics.StreamingJob{data}}, func(ctx context.Context, result streamanalytics.StreamingJobListResult) (streamanalytics.StreamingJobListResult, error) {
		return streamanalytics.StreamingJobListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "").AnyTimes().Return(result, nil)
	return s
}

func TestStreamAnalyticsStreamingJobs(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureStreamanalyticsStreamingJobsGenerator{}), createStreamingJobsMock, azure_client.TestOptions{})
}
