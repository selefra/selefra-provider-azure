package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureSqlDatabasesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureSqlDatabasesGenerator{}

func (x *TableAzureSqlDatabasesGenerator) GetTableName() string {
	return "azure_sql_databases"
}

func (x *TableAzureSqlDatabasesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureSqlDatabasesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureSqlDatabasesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureSqlDatabasesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().SQL.Databases

			server := task.ParentRawResult.(sql.Server)
			resourceDetails, err := azure_client.ParseResourceID(*server.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			response, err := svc.ListByServer(ctx, resourceDetails.ResourceGroup, *server.Name)

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

func (x *TableAzureSqlDatabasesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureSqlDatabasesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("managed_by").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_database_deletion_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("SourceDatabaseDeletionDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("catalog_collation").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("earliest_restore_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("EarliestRestoreDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sql_server_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sample_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_database_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceDatabaseID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_log_size_bytes").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_sku").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_size_bytes").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("license_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resumed_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("ResumedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("elastic_pool_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ElasticPoolID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DatabaseID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_service_objective_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requested_service_objective_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failover_group_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FailoverGroupID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restorable_dropped_database_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RestorableDroppedDatabaseID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_configuration_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MaintenanceConfigurationID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sku").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("paused_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("PausedDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("collation").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_date").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("CreationDate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("restore_point_in_time").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("RestorePointInTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recovery_services_recovery_point_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RecoveryServicesRecoveryPointID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone_redundant").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("read_scale").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("high_availability_replica_count").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_pause_delay").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_capacity").ColumnType(schema.ColumnTypeFloat).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recoverable_database_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RecoverableDatabaseID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_mode").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_secondary_location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("long_term_retention_backup_resource_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LongTermRetentionBackupResourceID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_account_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_sql_servers_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_sql_servers.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAzureSqlDatabasesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAzureSqlDatabaseVulnerabilityAssessmentScansGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlBackupLongTermRetentionPoliciesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlDatabaseThreatDetectionPoliciesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlTransparentDataEncryptionsGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlDatabaseBlobAuditingPoliciesGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureSqlDatabaseVulnerabilityAssessmentsGenerator{}),
	}
}
