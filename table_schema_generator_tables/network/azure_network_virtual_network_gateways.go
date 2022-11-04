package network

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureNetworkVirtualNetworkGatewaysGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureNetworkVirtualNetworkGatewaysGenerator{}

func (x *TableAzureNetworkVirtualNetworkGatewaysGenerator) GetTableName() string {
	return "azure_network_virtual_network_gateways"
}

func (x *TableAzureNetworkVirtualNetworkGatewaysGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureNetworkVirtualNetworkGatewaysGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureNetworkVirtualNetworkGatewaysGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureNetworkVirtualNetworkGatewaysGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().Network.VirtualNetworkGateways

			network := task.ParentRawResult.(network.VirtualNetwork)
			resourceDetails, err := azure_client.ParseResourceID(*network.ID)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			response, err := svc.List(ctx, resourceDetails.ResourceGroup)

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

func (x *TableAzureNetworkVirtualNetworkGatewaysGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableAzureNetworkVirtualNetworkGatewaysGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("provisioning_state").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_configurations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("IPConfigurations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpn_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sku").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_virtual_network_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_bgp").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inbound_dns_forwarding_endpoint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InboundDNSForwardingEndpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("azure_network_virtual_networks_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to azure_network_virtual_networks.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gateway_default_site").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpn_gateway_generation").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bgp_settings").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("extended_location").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gateway_type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_private_ip_address").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnablePrivateIPAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpn_client_configuration").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_routes").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_guid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ResourceGUID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active_active").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_dns_forwarding").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableDNSForwarding")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("v_net_extended_location_resource_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VNetExtendedLocationResourceID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableAzureNetworkVirtualNetworkGatewaysGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableAzureNetworkVirtualNetworkGatewayConnectionsGenerator{}),
	}
}
