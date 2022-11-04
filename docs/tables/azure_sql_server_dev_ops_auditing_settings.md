# Table: azure_sql_server_dev_ops_auditing_settings

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| storage_endpoint | string | X | √ |  | 
| storage_account_access_key | string | X | √ |  | 
| storage_account_subscription_id | string | X | √ |  | 
| id | string | √ | √ |  | 
| subscription_id | string | X | √ |  | 
| system_data | json | X | √ |  | 
| is_azure_monitor_target_enabled | bool | X | √ |  | 
| state | string | X | √ |  | 
| azure_sql_servers_selefra_id | string | X | X | fk to azure_sql_servers.selefra_id | 
| sql_server_id | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


