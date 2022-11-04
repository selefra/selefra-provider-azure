package monitor

import (
	"context"
	"regexp"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

type regexMatcher struct {
	re *regexp.Regexp
}

func (m regexMatcher) Matches(x interface{}) bool {
	s, ok := x.(string)
	if !ok {
		return false
	}
	return m.re.MatchString(s)
}
func (m regexMatcher) String() string {
	return m.re.String()
}

func createActivityLogsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMonitorActivityLogsClient(ctrl)
	s := services.Services{
		Monitor: services.MonitorClient{
			ActivityLogs: mockClient,
		},
	}

	data := insights.EventData{}
	require.Nil(t, faker.FakeObject(&data))

	result := insights.NewEventDataCollectionPage(insights.EventDataCollection{Value: &[]insights.EventData{data}}, func(ctx context.Context, result insights.EventDataCollection) (insights.EventDataCollection, error) {
		return insights.EventDataCollection{}, nil
	})

	filterRe := regexp.MustCompile(`eventTimestamp ge '\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d(\.\d+)Z' and eventTimestamp le '\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d(\.\d+)Z'`)
	mockClient.EXPECT().List(gomock.Any(), regexMatcher{filterRe}, "").AnyTimes().Return(result, nil)
	return s
}

func TestMonitorActivityLogs(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureMonitorActivityLogsGenerator{}), createActivityLogsMock, azure_client.TestOptions{})
}
