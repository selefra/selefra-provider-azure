package azure_client

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/test_helper"
	"github.com/spf13/viper"

	"github.com/selefra/selefra-provider-azure/azure_client/services"
)

type TestOptions struct{}

func MockTestHelper(t *testing.T, table *schema.Table, createServices func(t *testing.T, ctrl *gomock.Controller) services.Services, _ TestOptions) {
	testProvider := newTestProvider(t, table, createServices)
	config := "test : test"
	test_helper.RunProviderPullTables(testProvider, config, "./", "*")
}

func newTestProvider(t *testing.T, table *schema.Table, createServices func(t *testing.T, ctrl *gomock.Controller) services.Services) *provider.Provider {
	return &provider.Provider{
		Name:      "azure",
		Version:   "v0.0.1",
		TableList: []*schema.Table{table},
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				svc := createServices(t, gomock.NewController(t))
				return []any{
					&Client{
						subscriptions:  []string{"testSubscription"},
						SubscriptionId: "testSubscription",
						services: map[string]*services.Services{
							"testSubscription": &svc,
						},
					},
				}, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return ``
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
				"N/A",
				"not_supported",
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
