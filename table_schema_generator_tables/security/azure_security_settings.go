package security

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureSecuritySettingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureSecuritySettingsGenerator{}

func (x *TableAzureSecuritySettingsGenerator) GetTableName() string {
	return "azure_security_settings"
}

func (x *TableAzureSecuritySettingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureSecuritySettingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureSecuritySettingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureSecuritySettingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().Security.Settings

			response, err := svc.List(ctx)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			for response.NotDone() {
				for _, item := range response.Values() {
					if v, ok := item.AsSetting(); ok {
						resultChannel <- v
					} else if v, ok := item.AsDataExportSettings(); ok {
						resultChannel <- v
					} else if v, ok := item.AsAlertSyncSettings(); ok {
						resultChannel <- v
					} else {
						maybeError := errors.WithStack(fmt.Errorf("unexpected BasicSetting: %#v", item))
						return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
					}
				}
				if err := response.NextWithContext(ctx); err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
			}
			return nil
		},
	}
}

func (x *TableAzureSecuritySettingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureSecuritySettingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAzureSecuritySettingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
