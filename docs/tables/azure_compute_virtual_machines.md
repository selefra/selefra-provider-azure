# Table: azure_compute_virtual_machines

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| virtual_machine_scale_set | json | X | √ |  | 
| extended_location | json | X | √ |  | 
| name | string | X | √ |  | 
| additional_capabilities | json | X | √ |  | 
| eviction_policy | string | X | √ |  | 
| license_type | string | X | √ |  | 
| vm_id | string | X | √ |  | 
| zones | string_array | X | √ |  | 
| plan | json | X | √ |  | 
| priority | string | X | √ |  | 
| host_group | json | X | √ |  | 
| platform_fault_domain | int | X | √ |  | 
| type | string | X | √ |  | 
| network_profile | json | X | √ |  | 
| proximity_placement_group | json | X | √ |  | 
| location | string | X | √ |  | 
| security_profile | json | X | √ |  | 
| resources | json | X | √ |  | 
| id | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| instance_view | json | X | √ |  | 
| hardware_profile | json | X | √ |  | 
| billing_profile | json | X | √ |  | 
| host | json | X | √ |  | 
| identity | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| availability_set | json | X | √ |  | 
| extensions_time_budget | string | X | √ |  | 
| storage_profile | json | X | √ |  | 
| diagnostics_profile | json | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| scheduled_events_profile | json | X | √ |  | 
| user_data | string | X | √ |  | 
| tags | json | X | √ |  | 
| os_profile | json | X | √ |  | 


