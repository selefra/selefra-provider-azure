package network

import (
	"context"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureNetworkPublicIpAddressesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureNetworkPublicIpAddressesGenerator{}

func (x *TableAzureNetworkPublicIpAddressesGenerator) GetTableName() string {
	return "azure_network_public_ip_addresses"
}

func (x *TableAzureNetworkPublicIpAddressesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureNetworkPublicIpAddressesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureNetworkPublicIpAddressesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureNetworkPublicIpAddressesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().Network.PublicIPAddresses

			response, err := svc.ListAll(ctx)

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

func (x *TableAzureNetworkPublicIpAddressesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureNetworkPublicIpAddressesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("sku").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IPAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nat_gateway").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provisioning_state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("linked_public_ip_address").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LinkedPublicIPAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("extended_location").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_ip_allocation_method").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PublicIPAllocationMethod")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_ip_address_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PublicIPAddressVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("IPConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dns_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DNSSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("IPTags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("idle_timeout_in_minutes").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_guid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceGUID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_public_ip_address").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ServicePublicIPAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("migration_phase").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zones").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ddos_settings").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_ip_prefix").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PublicIPPrefix")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableAzureNetworkPublicIpAddressesGenerator) GetSubTables() []*schema.Table {
	return nil
}
