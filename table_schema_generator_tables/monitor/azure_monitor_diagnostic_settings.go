package monitor

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2021-07-01-preview/insights"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-10-01/resources"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureMonitorDiagnosticSettingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureMonitorDiagnosticSettingsGenerator{}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetTableName() string {
	return "azure_monitor_diagnostic_settings"
}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().Monitor.DiagnosticSettings

			resource := task.ParentRawResult.(resources.GenericResourceExpanded)
			response, err := svc.List(ctx, *resource.ID)
			if err != nil {
				if isResourceTypeNotSupported(err) {
					return nil
				}
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			if response.Value == nil {
				return nil
			}
			for _, v := range *response.Value {
				resultChannel <- diagnosticSettingResource{
					DiagnosticSettingsResource: v,
					ResourceURI:                *resource.ID,
				}
			}
			return nil
		},
	}
}

type diagnosticSettingResource struct {
	insights.DiagnosticSettingsResource
	ResourceURI string
}

func isResourceTypeNotSupported(err error) bool {
	var azureErr *azure.RequestError
	if errors.As(err, &azureErr) {
		return azureErr.ServiceError != nil && azureErr.ServiceError.Code == "ResourceTypeNotSupported"
	}
	return false
}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("metrics").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("workspace_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WorkspaceID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("monitor_resource_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_bus_rule_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceBusRuleID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_hub_authorization_rule_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EventHubAuthorizationRuleID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_analytics_destination_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceURI")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_account_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StorageAccountID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_hub_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logs").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_monitor_resources_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_monitor_resources.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAzureMonitorDiagnosticSettingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
