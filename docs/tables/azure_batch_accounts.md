# Table: azure_batch_accounts

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subscription_id | string | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| low_priority_core_quota | int | X | √ |  | 
| dedicated_core_quota_per_vm_family_enforced | bool | X | √ |  | 
| identity | json | X | √ |  | 
| name | string | X | √ |  | 
| tags | json | X | √ |  | 
| private_endpoint_connections | json | X | √ |  | 
| auto_storage | json | X | √ |  | 
| dedicated_core_quota_per_vm_family | json | X | √ |  | 
| allowed_authentication_modes | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| pool_allocation_mode | string | X | √ |  | 
| public_network_access | string | X | √ |  | 
| dedicated_core_quota | int | X | √ |  | 
| active_job_and_job_schedule_quota | int | X | √ |  | 
| id | string | √ | √ |  | 
| account_endpoint | string | X | √ |  | 
| key_vault_reference | json | X | √ |  | 
| encryption | json | X | √ |  | 
| pool_quota | int | X | √ |  | 
| type | string | X | √ |  | 
| location | string | X | √ |  | 


