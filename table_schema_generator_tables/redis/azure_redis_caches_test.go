package redis

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/redis/mgmt/2020-12-01/redis"
	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/azure_client/mocks"
	"github.com/selefra/selefra-provider-azure/azure_client/services"
	"github.com/selefra/selefra-provider-azure/faker"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/stretchr/testify/require"
)

func createCachesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockRedisCachesClient(ctrl)
	s := services.Services{
		Redis: services.RedisClient{
			Caches: mockClient,
		},
	}

	data := redis.ResourceType{}
	require.Nil(t, faker.FakeObject(&data))

	result := redis.NewListResultPage(redis.ListResult{Value: &[]redis.ResourceType{data}}, func(ctx context.Context, result redis.ListResult) (redis.ListResult, error) {
		return redis.ListResult{}, nil
	})

	mockClient.EXPECT().ListBySubscription(gomock.Any()).AnyTimes().Return(result, nil)
	return s
}

func TestRedisCaches(t *testing.T) {
	azure_client.MockTestHelper(t, table_schema_generator.GenTableSchema(&TableAzureRedisCachesGenerator{}), createCachesMock, azure_client.TestOptions{})
}
