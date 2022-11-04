# Table: azure_compute_virtual_machine_scale_sets

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| sku | json | X | √ |  | 
| type | string | X | √ |  | 
| additional_capabilities | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| single_placement_group | bool | X | √ |  | 
| automatic_repairs_policy | json | X | √ |  | 
| location | string | X | √ |  | 
| tags | json | X | √ |  | 
| zone_balance | bool | X | √ |  | 
| extended_location | json | X | √ |  | 
| host_group | json | X | √ |  | 
| identity | json | X | √ |  | 
| upgrade_policy | json | X | √ |  | 
| virtual_machine_profile | json | X | √ |  | 
| proximity_placement_group | json | X | √ |  | 
| orchestration_mode | string | X | √ |  | 
| unique_id | string | X | √ |  | 
| platform_fault_domain_count | int | X | √ |  | 
| zones | string_array | X | √ |  | 
| name | string | X | √ |  | 
| plan | json | X | √ |  | 
| scale_in_policy | json | X | √ |  | 
| id | string | √ | √ |  | 
| do_not_run_extensions_on_overprovisioned_vms | bool | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| overprovision | bool | X | √ |  | 


