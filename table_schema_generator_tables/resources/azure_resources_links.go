package resources

import (
	"context"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureResourcesLinksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureResourcesLinksGenerator{}

func (x *TableAzureResourcesLinksGenerator) GetTableName() string {
	return "azure_resources_links"
}

func (x *TableAzureResourcesLinksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureResourcesLinksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureResourcesLinksGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureResourcesLinksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().Resources.Links

			response, err := svc.ListAtSubscription(ctx, "")

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

func (x *TableAzureResourcesLinksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureResourcesLinksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_source_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.SourceID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_target_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.TargetID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_notes").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.Notes")).Build(),
	}
}

func (x *TableAzureResourcesLinksGenerator) GetSubTables() []*schema.Table {
	return nil
}
