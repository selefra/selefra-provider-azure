package provider

import (
	"github.com/selefra/selefra-provider-azure/table_schema_generator"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/authorization"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/batch"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/cdn"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/compute"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/container"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/cosmosdb"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/datalake"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/eventhub"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/frontdoor"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/iothub"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/keyvault"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/logic"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/mariadb"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/monitor"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/mysql"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/network"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/postgresql"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/redis"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/resources"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/search"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/security"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/servicebus"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/sql"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/storage"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/streamanalytics"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/subscriptions"
	"github.com/selefra/selefra-provider-azure/table_schema_generator_tables/web"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&iothub.TableAzureIothubDevicesGenerator{}),
		table_schema_generator.GenTableSchema(&monitor.TableAzureMonitorLogProfilesGenerator{}),
		table_schema_generator.GenTableSchema(&monitor.TableAzureMonitorActivityLogsGenerator{}),
		table_schema_generator.GenTableSchema(&monitor.TableAzureMonitorResourcesGenerator{}),
		table_schema_generator.GenTableSchema(&monitor.TableAzureMonitorActivityLogAlertsGenerator{}),
		table_schema_generator.GenTableSchema(&redis.TableAzureRedisCachesGenerator{}),
		table_schema_generator.GenTableSchema(&subscriptions.TableAzureSubscriptionsLocationsGenerator{}),
		table_schema_generator.GenTableSchema(&subscriptions.TableAzureSubscriptionsGenerator{}),
		table_schema_generator.GenTableSchema(&subscriptions.TableAzureSubscriptionsTenantsGenerator{}),
		table_schema_generator.GenTableSchema(&resources.TableAzureResourcesPolicyAssignmentsGenerator{}),
		table_schema_generator.GenTableSchema(&resources.TableAzureResourcesGroupsGenerator{}),
		table_schema_generator.GenTableSchema(&resources.TableAzureResourcesLinksGenerator{}),
		table_schema_generator.GenTableSchema(&storage.TableAzureStorageAccountsGenerator{}),
		table_schema_generator.GenTableSchema(&authorization.TableAzureAuthorizationRoleAssignmentsGenerator{}),
		table_schema_generator.GenTableSchema(&authorization.TableAzureAuthorizationRoleDefinitionsGenerator{}),
		table_schema_generator.GenTableSchema(&logic.TableAzureLogicWorkflowsGenerator{}),
		table_schema_generator.GenTableSchema(&mariadb.TableAzureMariadbServersGenerator{}),
		table_schema_generator.GenTableSchema(&servicebus.TableAzureServicebusNamespacesGenerator{}),
		table_schema_generator.GenTableSchema(&frontdoor.TableAzureFrontdoorDoorsGenerator{}),
		table_schema_generator.GenTableSchema(&search.TableAzureSearchServicesGenerator{}),
		table_schema_generator.GenTableSchema(&cosmosdb.TableAzureCosmosdbAccountsGenerator{}),
		table_schema_generator.GenTableSchema(&eventhub.TableAzureEventhubNamespacesGenerator{}),
		table_schema_generator.GenTableSchema(&datalake.TableAzureDatalakeAnalyticsAccountsGenerator{}),
		table_schema_generator.GenTableSchema(&datalake.TableAzureDatalakeStoreAccountsGenerator{}),
		table_schema_generator.GenTableSchema(&keyvault.TableAzureKeyvaultManagedHsmsGenerator{}),
		table_schema_generator.GenTableSchema(&keyvault.TableAzureKeyvaultVaultsGenerator{}),
		table_schema_generator.GenTableSchema(&postgresql.TableAzurePostgresqlServersGenerator{}),
		table_schema_generator.GenTableSchema(&cdn.TableAzureCdnProfilesGenerator{}),
		table_schema_generator.GenTableSchema(&container.TableAzureContainerManagedClustersGenerator{}),
		table_schema_generator.GenTableSchema(&container.TableAzureContainerRegistriesGenerator{}),
		table_schema_generator.GenTableSchema(&sql.TableAzureSqlManagedInstancesGenerator{}),
		table_schema_generator.GenTableSchema(&sql.TableAzureSqlServersGenerator{}),
		table_schema_generator.GenTableSchema(&security.TableAzureSecurityAssessmentsGenerator{}),
		table_schema_generator.GenTableSchema(&security.TableAzureSecurityContactsGenerator{}),
		table_schema_generator.GenTableSchema(&security.TableAzureSecurityJitNetworkAccessPoliciesGenerator{}),
		table_schema_generator.GenTableSchema(&security.TableAzureSecurityPricingsGenerator{}),
		table_schema_generator.GenTableSchema(&security.TableAzureSecuritySettingsGenerator{}),
		table_schema_generator.GenTableSchema(&security.TableAzureSecurityAutoProvisioningSettingsGenerator{}),
		table_schema_generator.GenTableSchema(&mysql.TableAzureMysqlServersGenerator{}),
		table_schema_generator.GenTableSchema(&network.TableAzureNetworkRouteFiltersGenerator{}),
		table_schema_generator.GenTableSchema(&network.TableAzureNetworkExpressRouteGatewaysGenerator{}),
		table_schema_generator.GenTableSchema(&network.TableAzureNetworkInterfacesGenerator{}),
		table_schema_generator.GenTableSchema(&network.TableAzureNetworkSecurityGroupsGenerator{}),
		table_schema_generator.GenTableSchema(&network.TableAzureNetworkPublicIpAddressesGenerator{}),
		table_schema_generator.GenTableSchema(&network.TableAzureNetworkRouteTablesGenerator{}),
		table_schema_generator.GenTableSchema(&network.TableAzureNetworkWatchersGenerator{}),
		table_schema_generator.GenTableSchema(&network.TableAzureNetworkVirtualNetworksGenerator{}),
		table_schema_generator.GenTableSchema(&network.TableAzureNetworkExpressRoutePortsGenerator{}),
		table_schema_generator.GenTableSchema(&network.TableAzureNetworkExpressRouteCircuitsGenerator{}),
		table_schema_generator.GenTableSchema(&streamanalytics.TableAzureStreamanalyticsStreamingJobsGenerator{}),
		table_schema_generator.GenTableSchema(&web.TableAzureWebAppsGenerator{}),
		table_schema_generator.GenTableSchema(&batch.TableAzureBatchAccountsGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableAzureComputeVirtualMachineScaleSetsGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableAzureComputeDisksGenerator{}),
		table_schema_generator.GenTableSchema(&compute.TableAzureComputeVirtualMachinesGenerator{}),
	}
}