package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureSqlManagedDatabasesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureSqlManagedDatabasesGenerator{}

func (x *TableAzureSqlManagedDatabasesGenerator) GetTableName() string {
	return "azure_sql_managed_databases"
}

func (x *TableAzureSqlManagedDatabasesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureSqlManagedDatabasesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureSqlManagedDatabasesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureSqlManagedDatabasesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().SQL.ManagedDatabases

			instance := task.ParentRawResult.(sql.ManagedInstance)
			resourceDetails, err := azure_client.ParseResourceID(*instance.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			response, err := svc.ListByInstance(ctx, resourceDetails.ResourceGroup, *instance.Name)

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

func (x *TableAzureSqlManagedDatabasesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureSqlManagedDatabasesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("source_database_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceDatabaseID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restorable_dropped_database_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RestorableDroppedDatabaseID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("long_term_retention_backup_resource_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LongTermRetentionBackupResourceID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("earliest_restore_point").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("EarliestRestorePoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recoverable_database_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RecoverableDatabaseID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_mode").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_container_sas_token").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failover_group_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FailoverGroupID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_backup_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sql_managed_instance_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restore_point_in_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("RestorePointInTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_secondary_location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("catalog_collation").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_container_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StorageContainerURI")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_complete_restore").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("collation").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("CreationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_sql_managed_instances_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_sql_managed_instances.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAzureSqlManagedDatabasesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAzureSqlManagedDatabaseVulnerabilityAssessmentsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlManagedDatabaseVulnerabilityAssessmentScansGenerator{}),
	}
}
