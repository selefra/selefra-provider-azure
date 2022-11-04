# Table: azure_iothub_devices

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| properties_disable_local_auth | bool | X | √ |  | 
| properties_enable_data_residency | bool | X | √ |  | 
| id | string | √ | √ |  | 
| properties_messaging_endpoints | json | X | √ |  | 
| etag | string | X | √ |  | 
| properties_authorization_policies | json | X | √ |  | 
| properties_ip_filter_rules | json | X | √ |  | 
| properties_min_tls_version | string | X | √ |  | 
| properties_cloud_to_device | json | X | √ |  | 
| type | string | X | √ |  | 
| subscription_id | string | X | √ |  | 
| properties_disable_device_sas | bool | X | √ |  | 
| properties_allowed_fqdn_list | string_array | X | √ |  | 
| properties_provisioning_state | string | X | √ |  | 
| properties_private_endpoint_connections | json | X | √ |  | 
| properties_enable_file_upload_notifications | bool | X | √ |  | 
| properties_comments | string | X | √ |  | 
| location | string | X | √ |  | 
| identity | json | X | √ |  | 
| properties_restrict_outbound_network_access | bool | X | √ |  | 
| properties_host_name | string | X | √ |  | 
| properties_event_hub_endpoints | json | X | √ |  | 
| properties_locations | json | X | √ |  | 
| properties_disable_module_sas | bool | X | √ |  | 
| properties_storage_endpoints | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| properties_public_network_access | string | X | √ |  | 
| properties_network_rule_sets | json | X | √ |  | 
| properties_routing | json | X | √ |  | 
| name | string | X | √ |  | 
| tags | json | X | √ |  | 
| properties_state | string | X | √ |  | 
| properties_features | string | X | √ |  | 
| sku | json | X | √ |  | 
| system_data | json | X | √ |  | 


