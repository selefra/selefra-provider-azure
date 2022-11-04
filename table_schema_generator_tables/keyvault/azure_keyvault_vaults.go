package keyvault

import (
	"context"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureKeyvaultVaultsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureKeyvaultVaultsGenerator{}

func (x *TableAzureKeyvaultVaultsGenerator) GetTableName() string {
	return "azure_keyvault_vaults"
}

func (x *TableAzureKeyvaultVaultsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureKeyvaultVaultsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureKeyvaultVaultsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureKeyvaultVaultsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().KeyVault.Vaults

			maxResults := int32(1000)
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

func (x *TableAzureKeyvaultVaultsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureKeyvaultVaultsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("properties_vault_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.VaultURI")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_enabled_for_deployment").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Properties.EnabledForDeployment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_soft_delete_retention_in_days").ColumnType(schema.ColumnTypeInt).
			Extractor(column_value_extractor.StructSelector("Properties.SoftDeleteRetentionInDays")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_enable_soft_delete").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Properties.EnableSoftDelete")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_tenant_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.TenantID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_enabled_for_disk_encryption").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Properties.EnabledForDiskEncryption")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_enabled_for_template_deployment").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Properties.EnabledForTemplateDeployment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_create_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.CreateMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_enable_purge_protection").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Properties.EnablePurgeProtection")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_network_acls").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties.NetworkAcls")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_private_endpoint_connections").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties.PrivateEndpointConnections")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_sku").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties.Sku")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_access_policies").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties.AccessPolicies")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_enable_rbac_authorization").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Properties.EnableRbacAuthorization")).Build(),
	}
}

func (x *TableAzureKeyvaultVaultsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAzureKeyvaultKeysGenerator{}),
		table_schema_generator.GenTableSchema(&TableAzureKeyvaultSecretsGenerator{}),
	}
}
