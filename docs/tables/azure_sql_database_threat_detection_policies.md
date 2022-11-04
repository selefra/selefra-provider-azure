# Table: azure_sql_database_threat_detection_policies

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| storage_endpoint | string | X | √ |  | 
| use_server_default | string | X | √ |  | 
| type | string | X | √ |  | 
| azure_sql_databases_selefra_id | string | X | X | fk to azure_sql_databases.selefra_id | 
| state | string | X | √ |  | 
| email_addresses | string | X | √ |  | 
| email_account_admins | string | X | √ |  | 
| name | string | X | √ |  | 
| disabled_alerts | string | X | √ |  | 
| retention_days | int | X | √ |  | 
| subscription_id | string | X | √ |  | 
| sql_database_id | string | X | √ |  | 
| location | string | X | √ |  | 
| kind | string | X | √ |  | 
| storage_account_access_key | string | X | √ |  | 
| id | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


