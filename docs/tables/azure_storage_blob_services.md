# Table: azure_storage_blob_services

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | √ | √ |  | 
| restore_policy | json | X | √ |  | 
| automatic_snapshot_policy_enabled | bool | X | √ |  | 
| default_service_version | string | X | √ |  | 
| delete_retention_policy | json | X | √ |  | 
| is_versioning_enabled | bool | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| cors | json | X | √ |  | 
| storage_account_id | string | X | √ |  | 
| change_feed | json | X | √ |  | 
| container_delete_retention_policy | json | X | √ |  | 
| last_access_time_tracking_policy | json | X | √ |  | 
| sku | json | X | √ |  | 
| type | string | X | √ |  | 
| azure_storage_accounts_selefra_id | string | X | X | fk to azure_storage_accounts.selefra_id | 
| subscription_id | string | X | √ |  | 


