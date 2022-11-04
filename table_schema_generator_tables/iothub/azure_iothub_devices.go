package iothub

import (
	"context"

	"github.com/selefra/selefra-provider-azure/azure_client"
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableAzureIothubDevicesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableAzureIothubDevicesGenerator{}

func (x *TableAzureIothubDevicesGenerator) GetTableName() string {
	return "azure_iothub_devices"
}

func (x *TableAzureIothubDevicesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableAzureIothubDevicesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableAzureIothubDevicesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableAzureIothubDevicesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			svc := client.(*azure_client.Client).AzureServices().IotHub.Devices

			response, err := svc.ListBySubscription(ctx)

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

func (x *TableAzureIothubDevicesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return azure_client.ExpandSubscription()
}

func (x *TableAzureIothubDevicesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("properties_disable_local_auth").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Properties.DisableLocalAuth")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_enable_data_residency").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Properties.EnableDataResidency")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_messaging_endpoints").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties.MessagingEndpoints")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_authorization_policies").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties.AuthorizationPolicies")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_ip_filter_rules").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties.IPFilterRules")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_min_tls_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.MinTLSVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_cloud_to_device").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties.CloudToDevice")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscription_id").ColumnType(schema.ColumnTypeString).
			Extractor(azure_client.ExtractorAzureSubscription()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_disable_device_sas").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Properties.DisableDeviceSAS")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_allowed_fqdn_list").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Properties.AllowedFqdnList")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_provisioning_state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.ProvisioningState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_private_endpoint_connections").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties.PrivateEndpointConnections")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_enable_file_upload_notifications").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Properties.EnableFileUploadNotifications")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_comments").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.Comments")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_restrict_outbound_network_access").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Properties.RestrictOutboundNetworkAccess")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_host_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.HostName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_event_hub_endpoints").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties.EventHubEndpoints")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_locations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties.Locations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_disable_module_sas").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Properties.DisableModuleSAS")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_storage_endpoints").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties.StorageEndpoints")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_public_network_access").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.PublicNetworkAccess")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_network_rule_sets").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties.NetworkRuleSets")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_routing").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Properties.Routing")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("properties_features").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Properties.Features")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sku").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("system_data").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableAzureIothubDevicesGenerator) GetSubTables() []*schema.Table {
	return nil
}
