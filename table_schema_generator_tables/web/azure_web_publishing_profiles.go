package web

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureWebPublishingProfilesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureWebPublishingProfilesGenerator{}

func (x *TableAzureWebPublishingProfilesGenerator) GetTableName() string {
	return "azure_web_publishing_profiles"
}

func (x *TableAzureWebPublishingProfilesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureWebPublishingProfilesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureWebPublishingProfilesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableAzureWebPublishingProfilesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().Web.PublishingProfiles

			site := task.ParentRawResult.(web.Site)
			response, err := svc.ListPublishingProfiles(ctx, *site.ResourceGroup, *site.Name)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			resultChannel <- response
			return nil
		},
	}
}

func (x *TableAzureWebPublishingProfilesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureWebPublishingProfilesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("user_pwd").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UserPWD")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_web_apps_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_web_apps.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_app_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("publish_url").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_name").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAzureWebPublishingProfilesGenerator) GetSubTables() []*schema.Table {
	return nil
}
