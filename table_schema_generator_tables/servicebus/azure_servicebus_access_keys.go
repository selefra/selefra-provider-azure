package servicebus

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureServicebusAccessKeysGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureServicebusAccessKeysGenerator{}

func (x *TableAzureServicebusAccessKeysGenerator) GetTableName() string {
	return "azure_servicebus_access_keys"
}

func (x *TableAzureServicebusAccessKeysGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureServicebusAccessKeysGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureServicebusAccessKeysGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureServicebusAccessKeysGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().Servicebus.AccessKeys

			namespace := task.ParentTask.ParentTask.ParentRawResult.(servicebus.SBNamespace)
			topic := task.ParentTask.ParentRawResult.(servicebus.SBTopic)
			rule := task.ParentRawResult.(servicebus.SBAuthorizationRule)
			resourceDetails, err := azure_client.ParseResourceID(*rule.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			response, err := svc.ListKeys(ctx, resourceDetails.ResourceGroup, *namespace.Name, *topic.Name, *rule.Name)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- response
			return nil
		},
	}
}

func (x *TableAzureServicebusAccessKeysGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureServicebusAccessKeysGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("primary_connection_string").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_connection_string").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("servicebus_authorization_rule_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("primary_key").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_key").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_servicebus_authorization_rules_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_servicebus_authorization_rules.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alias_primary_connection_string").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alias_secondary_connection_string").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAzureServicebusAccessKeysGenerator) GetSubTables() []*schema.Table {
	return nil
}
