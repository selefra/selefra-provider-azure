package sql

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureSqlTransparentDataEncryptionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureSqlTransparentDataEncryptionsGenerator{}

func (x *TableAzureSqlTransparentDataEncryptionsGenerator) GetTableName() string {
	return "azure_sql_transparent_data_encryptions"
}

func (x *TableAzureSqlTransparentDataEncryptionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureSqlTransparentDataEncryptionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureSqlTransparentDataEncryptionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureSqlTransparentDataEncryptionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().SQL.TransparentDataEncryptions

			server := task.ParentTask.ParentRawResult.(sql.Server)
			database := task.ParentRawResult.(sql.Database)
			resourceDetails, err := azure_client.ParseResourceID(*database.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			response, err := svc.Get(ctx, resourceDetails.ResourceGroup, *server.Name, *database.Name)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- response
			return nil
		},
	}
}

func (x *TableAzureSqlTransparentDataEncryptionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureSqlTransparentDataEncryptionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sql_database_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_sql_databases_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_sql_databases.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableAzureSqlTransparentDataEncryptionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
