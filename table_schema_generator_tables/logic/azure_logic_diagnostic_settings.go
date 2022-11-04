package logic

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/logic/mgmt/2019-05-01/logic"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureLogicDiagnosticSettingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureLogicDiagnosticSettingsGenerator{}

func (x *TableAzureLogicDiagnosticSettingsGenerator) GetTableName() string {
	return "azure_logic_diagnostic_settings"
}

func (x *TableAzureLogicDiagnosticSettingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureLogicDiagnosticSettingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureLogicDiagnosticSettingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureLogicDiagnosticSettingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().Logic.DiagnosticSettings

			workflow := task.ParentRawResult.(logic.Workflow)
			response, err := svc.List(ctx, *workflow.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			if response.Value == nil {
				return nil
			}
			resultChannel <- *response.Value

			return nil
		},
	}
}

func (x *TableAzureLogicDiagnosticSettingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureLogicDiagnosticSettingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logic_workflow_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metrics").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_account_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StorageAccountID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_hub_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("workspace_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WorkspaceID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_bus_rule_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceBusRuleID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_hub_authorization_rule_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EventHubAuthorizationRuleID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logs").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_analytics_destination_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_logic_workflows_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_logic_workflows.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAzureLogicDiagnosticSettingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
