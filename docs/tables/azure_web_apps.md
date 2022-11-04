# Table: azure_web_apps

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| is_xenon | bool | X | √ |  | 
| custom_domain_verification_id | string | X | √ |  | 
| daily_memory_time_quota | int | X | √ |  | 
| max_number_of_workers | int | X | √ |  | 
| slot_swap_status | json | X | √ |  | 
| storage_account_required | bool | X | √ |  | 
| subscription_id | string | X | √ |  | 
| reserved | bool | X | √ |  | 
| type | string | X | √ |  | 
| server_farm_id | string | X | √ |  | 
| traffic_manager_host_names | string_array | X | √ |  | 
| last_modified_time_utc | timestamp | X | √ |  | 
| scm_site_also_stopped | bool | X | √ |  | 
| client_affinity_enabled | bool | X | √ |  | 
| outbound_ip_addresses | string | X | √ |  | 
| cloning_info | json | X | √ |  | 
| name | string | X | √ |  | 
| enabled_host_names | string_array | X | √ |  | 
| host_name_ssl_states | json | X | √ |  | 
| location | string | X | √ |  | 
| https_only | bool | X | √ |  | 
| in_progress_operation_id | string | X | √ |  | 
| virtual_network_subnet_id | string | X | √ |  | 
| identity | json | X | √ |  | 
| state | string | X | √ |  | 
| client_cert_enabled | bool | X | √ |  | 
| resource_group | string | X | √ |  | 
| repository_site_name | string | X | √ |  | 
| usage_state | string | X | √ |  | 
| availability_state | string | X | √ |  | 
| target_swap_slot | string | X | √ |  | 
| hosting_environment_profile | json | X | √ |  | 
| client_cert_mode | string | X | √ |  | 
| host_names_disabled | bool | X | √ |  | 
| suspended_till | timestamp | X | √ |  | 
| host_names | string_array | X | √ |  | 
| enabled | bool | X | √ |  | 
| kind | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| default_host_name | string | X | √ |  | 
| redundancy_mode | string | X | √ |  | 
| container_size | int | X | √ |  | 
| is_default_container | bool | X | √ |  | 
| id | string | √ | √ |  | 
| hyper_v | bool | X | √ |  | 
| site_config | json | X | √ |  | 
| key_vault_reference_identity | string | X | √ |  | 
| tags | json | X | √ |  | 
| client_cert_exclusion_paths | string | X | √ |  | 
| possible_outbound_ip_addresses | string | X | √ |  | 


