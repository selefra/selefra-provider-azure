package subscriptions

import (
	"context"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureSubscriptionsTenantsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureSubscriptionsTenantsGenerator{}

func (x *TableAzureSubscriptionsTenantsGenerator) GetTableName() string {
	return "azure_subscriptions_tenants"
}

func (x *TableAzureSubscriptionsTenantsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureSubscriptionsTenantsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureSubscriptionsTenantsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureSubscriptionsTenantsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().Subscriptions.Tenants
			pager := svc.NewListPager(nil)
			for pager.More() {
				nextResult, err := pager.NextPage(ctx)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				for _, v := range nextResult.Value {
					resultChannel <- v
				}
			}
			return nil
		},
	}
}

func (x *TableAzureSubscriptionsTenantsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureSubscriptionsTenantsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("default_domain").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domains").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("country_code").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("country").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_branding_logo_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TenantBrandingLogoURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_category").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tenant_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TenantID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
	}
}

func (x *TableAzureSubscriptionsTenantsGenerator) GetSubTables() []*schema.Table {
	return nil
}
