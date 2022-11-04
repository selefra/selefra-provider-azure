# Table: azure_sql_database_blob_auditing_policies

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| retention_days | int | X | √ |  | 
| queue_delay_ms | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| sql_database_id | string | X | √ |  | 
| storage_endpoint | string | X | √ |  | 
| audit_actions_and_groups | string_array | X | √ |  | 
| is_storage_secondary_key_in_use | bool | X | √ |  | 
| is_azure_monitor_target_enabled | bool | X | √ |  | 
| id | string | √ | √ |  | 
| azure_sql_databases_selefra_id | string | X | X | fk to azure_sql_databases.selefra_id | 
| subscription_id | string | X | √ |  | 
| storage_account_access_key | string | X | √ |  | 
| storage_account_subscription_id | string | X | √ |  | 
| type | string | X | √ |  | 
| kind | string | X | √ |  | 
| state | string | X | √ |  | 
| name | string | X | √ |  | 


