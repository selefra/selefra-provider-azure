package sql

import (
	"context"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureSqlServersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureSqlServersGenerator{}

func (x *TableAzureSqlServersGenerator) GetTableName() string {
	return "azure_sql_servers"
}

func (x *TableAzureSqlServersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureSqlServersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureSqlServersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureSqlServersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().SQL.Servers

			response, err := svc.List(ctx)

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

func (x *TableAzureSqlServersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureSqlServersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("public_network_access").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("administrator_login").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fully_qualified_domain_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_endpoint_connections").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("administrator_login_password").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimal_tls_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MinimalTLSVersion")).Build(),
	}
}

func (x *TableAzureSqlServersGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAzureSqlServerAdminsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlServerVulnerabilityAssessmentsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlServerSecurityAlertPoliciesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlVirtualNetworkRulesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlDatabasesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlEncryptionProtectorsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlServerBlobAuditingPoliciesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlServerDevOpsAuditingSettingsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlFirewallRulesGenerator{}),
	}
}
