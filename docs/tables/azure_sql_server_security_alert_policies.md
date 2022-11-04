# Table: azure_sql_server_security_alert_policies

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| storage_account_access_key | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| state | string | X | √ |  | 
| email_addresses | string_array | X | √ |  | 
| email_account_admins | bool | X | √ |  | 
| storage_endpoint | string | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| sql_server_id | string | X | √ |  | 
| retention_days | int | X | √ |  | 
| azure_sql_servers_selefra_id | string | X | X | fk to azure_sql_servers.selefra_id | 
| subscription_id | string | X | √ |  | 
| disabled_alerts | string_array | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| type | string | X | √ |  | 


