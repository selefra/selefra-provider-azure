package postgresql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/postgresql/mgmt/postgresql"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzurePostgresqlFirewallRulesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzurePostgresqlFirewallRulesGenerator{}

func (x *TableAzurePostgresqlFirewallRulesGenerator) GetTableName() string {
	return "azure_postgresql_firewall_rules"
}

func (x *TableAzurePostgresqlFirewallRulesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzurePostgresqlFirewallRulesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzurePostgresqlFirewallRulesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzurePostgresqlFirewallRulesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().PostgreSQL.FirewallRules

			server := task.ParentRawResult.(postgresql.Server)
			resourceDetails, err := azure_client.ParseResourceID(*server.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			response, err := svc.ListByServer(ctx, resourceDetails.ResourceGroup, *server.Name)
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

func (x *TableAzurePostgresqlFirewallRulesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzurePostgresqlFirewallRulesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("postgresql_server_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_ip_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StartIPAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_ip_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EndIPAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_postgresql_servers_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_postgresql_servers.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAzurePostgresqlFirewallRulesGenerator) GetSubTables() []*schema.Table {
	return nil
}
