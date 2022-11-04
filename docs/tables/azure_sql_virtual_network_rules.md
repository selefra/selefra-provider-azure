# Table: azure_sql_virtual_network_rules

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| azure_sql_servers_selefra_id | string | X | X | fk to azure_sql_servers.selefra_id | 
| id | string | √ | √ |  | 
| sql_server_id | string | X | √ |  | 
| virtual_network_subnet_id | string | X | √ |  | 
| ignore_missing_vnet_service_endpoint | bool | X | √ |  | 
| state | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| subscription_id | string | X | √ |  | 


