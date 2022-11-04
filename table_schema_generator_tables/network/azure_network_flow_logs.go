package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureNetworkFlowLogsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureNetworkFlowLogsGenerator{}

func (x *TableAzureNetworkFlowLogsGenerator) GetTableName() string {
	return "azure_network_flow_logs"
}

func (x *TableAzureNetworkFlowLogsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureNetworkFlowLogsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureNetworkFlowLogsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureNetworkFlowLogsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().Network.FlowLogs

			watcher := task.ParentRawResult.(network.Watcher)
			resourceDetails, err := azure_client.ParseResourceID(*watcher.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			response, err := svc.List(ctx, resourceDetails.ResourceGroup, *watcher.Name)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			for response.NotDone() {
				resultChannel <- response.Values()
				if err := response.NextWithContext(ctx); err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
			}

			return nil
		},
	}
}

func (x *TableAzureNetworkFlowLogsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureNetworkFlowLogsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("network_watcher_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_resource_guid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TargetResourceGUID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retention_policy").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StorageID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("format").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_network_watchers_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_network_watchers.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_resource_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TargetResourceID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("flow_analytics_configuration").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provisioning_state").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAzureNetworkFlowLogsGenerator) GetSubTables() []*schema.Table {
	return nil
}
