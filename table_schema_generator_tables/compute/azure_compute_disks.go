package compute

import (
	"context"

	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
)

type TableAzureComputeDisksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureComputeDisksGenerator{}

func (x *TableAzureComputeDisksGenerator) GetTableName() string {
	return "azure_compute_disks"
}

func (x *TableAzureComputeDisksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureComputeDisksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureComputeDisksGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureComputeDisksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().Compute.Disks

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

func (x *TableAzureComputeDisksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureComputeDisksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("disk_state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_profile").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_size_bytes").ColumnType(schema.ColumnTypeBigInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zones").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supports_hibernation").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("purchase_plan").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_iops_read_write").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("DiskIOPSReadWrite")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_m_bps_read_write").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("share_info").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tier").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("managed_by").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("managed_by_extended").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sku").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_created").ColumnType(schema.ColumnTypeTimestamp).
			Extractor(azure_client.ExtractorAzureDateTime("TimeCreated")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_settings_collection").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_shares").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_size_gb").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("DiskSizeGB")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unique_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UniqueID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("extended_location").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("os_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_data").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hyper_v_generation").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_access_policy").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_m_bps_read_only").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_access_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DiskAccessID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bursting_enabled").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("property_updates_in_progress").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provisioning_state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_iops_read_only").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("DiskIOPSReadOnly")).Build(),
	}
}

func (x *TableAzureComputeDisksGenerator) GetSubTables() []*schema.Table {
	return nil
}
