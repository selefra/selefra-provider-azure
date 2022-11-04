package authorization

import (
	"context"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureAuthorizationRoleAssignmentsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureAuthorizationRoleAssignmentsGenerator{}

func (x *TableAzureAuthorizationRoleAssignmentsGenerator) GetTableName() string {
	return "azure_authorization_role_assignments"
}

func (x *TableAzureAuthorizationRoleAssignmentsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureAuthorizationRoleAssignmentsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureAuthorizationRoleAssignmentsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureAuthorizationRoleAssignmentsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().Authorization.RoleAssignments

			response, err := svc.List(ctx, "")

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

func (x *TableAzureAuthorizationRoleAssignmentsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureAuthorizationRoleAssignmentsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("properties_principal_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.PrincipalID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_scope").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.Scope")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_role_definition_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.RoleDefinitionID")).Build(),
	}
}

func (x *TableAzureAuthorizationRoleAssignmentsGenerator) GetSubTables() []*schema.Table {
	return nil
}
