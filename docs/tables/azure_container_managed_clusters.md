# Table: azure_container_managed_clusters

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| agent_pool_profiles | json | X | √ |  | 
| linux_profile | json | X | √ |  | 
| disable_local_accounts | bool | X | √ |  | 
| identity | json | X | √ |  | 
| http_proxy_config | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| power_state | json | X | √ |  | 
| fqdn_subdomain | string | X | √ |  | 
| network_profile | json | X | √ |  | 
| addon_profiles | json | X | √ |  | 
| enable_rbac | bool | X | √ |  | 
| api_server_access_profile | json | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| max_agent_pools | int | X | √ |  | 
| fqdn | string | X | √ |  | 
| private_fqdn | string | X | √ |  | 
| extended_location | json | X | √ |  | 
| node_resource_group | string | X | √ |  | 
| identity_profile | json | X | √ |  | 
| private_link_resources | json | X | √ |  | 
| sku | json | X | √ |  | 
| pod_identity_profile | json | X | √ |  | 
| auto_upgrade_profile | json | X | √ |  | 
| tags | json | X | √ |  | 
| kubernetes_version | string | X | √ |  | 
| azure_portal_fqdn | string | X | √ |  | 
| windows_profile | json | X | √ |  | 
| aad_profile | json | X | √ |  | 
| disk_encryption_set_id | string | X | √ |  | 
| type | string | X | √ |  | 
| location | string | X | √ |  | 
| dns_prefix | string | X | √ |  | 
| service_principal_profile | json | X | √ |  | 
| enable_pod_security_policy | bool | X | √ |  | 
| auto_scaler_profile | json | X | √ |  | 


