package mariadb

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/mariadb/mgmt/mariadb"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureMariadbConfigurationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureMariadbConfigurationsGenerator{}

func (x *TableAzureMariadbConfigurationsGenerator) GetTableName() string {
	return "azure_mariadb_configurations"
}

func (x *TableAzureMariadbConfigurationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureMariadbConfigurationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureMariadbConfigurationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureMariadbConfigurationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().MariaDB.Configurations

			server := task.ParentRawResult.(mariadb.Server)
			resourceDetails, err := azure_client.ParseResourceID(*server.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			response, err := svc.ListByServer(ctx, resourceDetails.ResourceGroup, *server.Name)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			if response.Value == nil {
				return nil
			}
			resultChannel <- *response.Value

			return nil
		},
	}
}

func (x *TableAzureMariadbConfigurationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureMariadbConfigurationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_mariadb_servers_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_mariadb_servers.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_values").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mariadb_server_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("value").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_value").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("data_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAzureMariadbConfigurationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
