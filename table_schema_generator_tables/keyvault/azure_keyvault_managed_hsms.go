package keyvault

import (
	"context"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureKeyvaultManagedHsmsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureKeyvaultManagedHsmsGenerator{}

func (x *TableAzureKeyvaultManagedHsmsGenerator) GetTableName() string {
	return "azure_keyvault_managed_hsms"
}

func (x *TableAzureKeyvaultManagedHsmsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureKeyvaultManagedHsmsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureKeyvaultManagedHsmsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureKeyvaultManagedHsmsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().KeyVault.ManagedHsms

			maxResults := int32(100)
			response, err := svc.ListBySubscription(ctx, &maxResults)

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

func (x *TableAzureKeyvaultManagedHsmsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureKeyvaultManagedHsmsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("properties_enable_purge_protection").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Properties.EnablePurgeProtection")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_enable_soft_delete").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Properties.EnableSoftDelete")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_provisioning_state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.ProvisioningState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_hsm_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.HsmURI")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_soft_delete_retention_in_days").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("Properties.SoftDeleteRetentionInDays")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_status_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.StatusMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_tenant_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.TenantID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_initial_admin_object_ids").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Properties.InitialAdminObjectIds")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_create_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.CreateMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sku").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAzureKeyvaultManagedHsmsGenerator) GetSubTables() []*schema.Table {
	return nil
}
