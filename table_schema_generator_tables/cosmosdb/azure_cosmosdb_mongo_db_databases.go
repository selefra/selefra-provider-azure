package cosmosdb

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/cosmos-db/mgmt/2020-04-01-preview/documentdb"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureCosmosdbMongoDbDatabasesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureCosmosdbMongoDbDatabasesGenerator{}

func (x *TableAzureCosmosdbMongoDbDatabasesGenerator) GetTableName() string {
	return "azure_cosmosdb_mongo_db_databases"
}

func (x *TableAzureCosmosdbMongoDbDatabasesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureCosmosdbMongoDbDatabasesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureCosmosdbMongoDbDatabasesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureCosmosdbMongoDbDatabasesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().CosmosDB.MongoDBDatabases

			account := task.ParentRawResult.(documentdb.DatabaseAccountGetResults)
			resource, err := azure_client.ParseResourceID(*account.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			response, err := svc.ListMongoDBDatabases(ctx, resource.ResourceGroup, *account.Name)
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

func (x *TableAzureCosmosdbMongoDbDatabasesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureCosmosdbMongoDbDatabasesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("azure_cosmosdb_accounts_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_cosmosdb_accounts.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("options").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cosmosdb_account_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAzureCosmosdbMongoDbDatabasesGenerator) GetSubTables() []*schema.Table {
	return nil
}
