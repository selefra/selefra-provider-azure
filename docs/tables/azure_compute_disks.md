# Table: azure_compute_disks

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| disk_state | string | X | √ |  | 
| security_profile | json | X | √ |  | 
| name | string | X | √ |  | 
| disk_size_bytes | int | X | √ |  | 
| encryption | json | X | √ |  | 
| type | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| zones | string_array | X | √ |  | 
| supports_hibernation | bool | X | √ |  | 
| id | string | √ | √ |  | 
| purchase_plan | json | X | √ |  | 
| disk_iops_read_write | int | X | √ |  | 
| disk_m_bps_read_write | int | X | √ |  | 
| share_info | json | X | √ |  | 
| tier | string | X | √ |  | 
| managed_by | string | X | √ |  | 
| managed_by_extended | string_array | X | √ |  | 
| sku | json | X | √ |  | 
| time_created | timestamp | X | √ |  | 
| encryption_settings_collection | json | X | √ |  | 
| max_shares | int | X | √ |  | 
| disk_size_gb | int | X | √ |  | 
| unique_id | string | X | √ |  | 
| location | string | X | √ |  | 
| extended_location | json | X | √ |  | 
| os_type | string | X | √ |  | 
| creation_data | json | X | √ |  | 
| hyper_v_generation | string | X | √ |  | 
| network_access_policy | string | X | √ |  | 
| disk_m_bps_read_only | int | X | √ |  | 
| disk_access_id | string | X | √ |  | 
| bursting_enabled | bool | X | √ |  | 
| property_updates_in_progress | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| disk_iops_read_only | int | X | √ |  | 


