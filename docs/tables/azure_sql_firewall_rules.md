# Table: azure_sql_firewall_rules

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subscription_id | string | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| azure_sql_servers_selefra_id | string | X | X | fk to azure_sql_servers.selefra_id | 
| sql_server_id | string | X | √ |  | 
| kind | string | X | √ |  | 
| location | string | X | √ |  | 
| start_ip_address | string | X | √ |  | 
| end_ip_address | string | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


