# Table: azure_sql_server_admins

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| administrator_type | string | X | √ |  | 
| sid | string | X | √ |  | 
| tenant_id | string | X | √ |  | 
| azure_ad_only_authentication | bool | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| subscription_id | string | X | √ |  | 
| sql_server_id | string | X | √ |  | 
| login | string | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| azure_sql_servers_selefra_id | string | X | X | fk to azure_sql_servers.selefra_id | 


