package sql

import (
	"context"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureSqlManagedInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureSqlManagedInstancesGenerator{}

func (x *TableAzureSqlManagedInstancesGenerator) GetTableName() string {
	return "azure_sql_managed_instances"
}

func (x *TableAzureSqlManagedInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureSqlManagedInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureSqlManagedInstancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureSqlManagedInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().SQL.ManagedInstances

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

func (x *TableAzureSqlManagedInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureSqlManagedInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("provisioning_state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("administrator_login_password").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("v_cores").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restore_point_in_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("RestorePointInTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("collation").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dns_zone_partner").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DNSZonePartner")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_endpoint_connections").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_redundant").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SubnetID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("license_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dns_zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DNSZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_data_endpoint_enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("administrator_login").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_pool_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstancePoolID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("minimal_tls_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MinimalTLSVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timezone_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TimezoneID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sku").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_account_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("managed_instance_create_mode").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fully_qualified_domain_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_size_in_gb").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("StorageSizeInGB")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_managed_instance_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceManagedInstanceID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("proxy_override").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_configuration_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MaintenanceConfigurationID")).Build(),
	}
}

func (x *TableAzureSqlManagedInstancesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAzureSqlManagedDatabasesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlManagedInstanceVulnerabilityAssessmentsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlManagedInstanceEncryptionProtectorsGenerator{}),
	}
}
