# Table: azure_sql_server_blob_auditing_policies

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| storage_endpoint | string | X | √ |  | 
| storage_account_subscription_id | string | X | √ |  | 
| is_azure_monitor_target_enabled | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| sql_server_id | string | X | √ |  | 
| queue_delay_ms | int | X | √ |  | 
| id | string | √ | √ |  | 
| azure_sql_servers_selefra_id | string | X | X | fk to azure_sql_servers.selefra_id | 
| audit_actions_and_groups | string_array | X | √ |  | 
| storage_account_access_key | string | X | √ |  | 
| is_storage_secondary_key_in_use | bool | X | √ |  | 
| subscription_id | string | X | √ |  | 
| retention_days | int | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| state | string | X | √ |  | 


