package azure_client

import (
	"context"
	"time"

	"github.com/Azure/go-autorest/autorest/date"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-utils/pkg/reflect_util"
)

func ExtractorAzureSubscription() schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		return client.(*Client).SubscriptionId, nil
	})
}

func ExtractorAzureDateTime(col string) schema.ColumnValueExtractor {
	return column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {
		if reflect_util.IsNil(result) {
			return nil, nil
		}
		value, err := column_value_extractor.StructSelector(col).Extract(ctx, clientMeta, client, task, row, column, result)
		if err != nil && err.HasError() {
			return nil, err
		}
		if reflect_util.IsNil(value) {
			return time.Time{}, nil
		}
		switch v := value.(type) {
		case *date.Time:
			return v.ToTime(), nil
		case *time.Time:
			return v, nil
		case date.Time:
			return v.ToTime(), nil
		case time.Time:
			return v, nil
		default:
			return nil, schema.NewDiagnostics().AddErrorMsg("Invalid type %T", v)
		}
	})
}
