package compute

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-03-01/compute"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureComputeInstanceViewsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureComputeInstanceViewsGenerator{}

func (x *TableAzureComputeInstanceViewsGenerator) GetTableName() string {
	return "azure_compute_instance_views"
}

func (x *TableAzureComputeInstanceViewsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureComputeInstanceViewsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureComputeInstanceViewsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureComputeInstanceViewsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().Compute.InstanceViews

			virtualMachine := task.ParentRawResult.(compute.VirtualMachine)
			resource, err := azure_client.ParseResourceID(*virtualMachine.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			response, err := svc.InstanceView(ctx, resource.ResourceGroup, *virtualMachine.Name)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- response
			return nil
		},
	}
}

func (x *TableAzureComputeInstanceViewsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureComputeInstanceViewsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("rdp_thumb_print").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_compute_virtual_machines_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_compute_virtual_machines.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("os_version").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("os_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("patch_status").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compute_virtual_machine_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("computer_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hyper_v_generation").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_redeploy_status").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vm_health").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VMHealth")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("boot_diagnostics").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("assigned_host").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("statuses").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_update_domain").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("platform_fault_domain").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vm_agent").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VMAgent")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disks").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("extensions").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
	}
}

func (x *TableAzureComputeInstanceViewsGenerator) GetSubTables() []*schema.Table {
	return nil
}
