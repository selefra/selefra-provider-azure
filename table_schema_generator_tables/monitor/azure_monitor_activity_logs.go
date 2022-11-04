package monitor

import (
	"context"
	"fmt"
	"time"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureMonitorActivityLogsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureMonitorActivityLogsGenerator{}

func (x *TableAzureMonitorActivityLogsGenerator) GetTableName() string {
	return "azure_monitor_activity_logs"
}

func (x *TableAzureMonitorActivityLogsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureMonitorActivityLogsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureMonitorActivityLogsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureMonitorActivityLogsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().Monitor.ActivityLogs

			const fetchWindow = 24 * time.Hour
			now := time.Now().UTC()
			past := now.Add(-fetchWindow)
			filter := fmt.Sprintf("eventTimestamp ge '%s' and eventTimestamp le '%s'", past.Format(time.RFC3339Nano), now.Format(time.RFC3339Nano))
			response, err := svc.List(ctx, filter, "")

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

func (x *TableAzureMonitorActivityLogsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureMonitorActivityLogsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("operation_name").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_type").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sub_status").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorization").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_data_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EventDataID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("correlation_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CorrelationID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("submission_timestamp").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("SubmissionTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TenantID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("category").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("level").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_group_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_provider_name").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operation_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OperationID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_timestamp").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("EventTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("claims").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("caller").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_name").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("http_request").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("HTTPRequest")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceID")).Build(),
	}
}

func (x *TableAzureMonitorActivityLogsGenerator) GetSubTables() []*schema.Table {
	return nil
}
