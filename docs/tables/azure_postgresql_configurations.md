# Table: azure_postgresql_configurations

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subscription_id | string | X | √ |  | 
| postgresql_server_id | string | X | √ |  | 
| description | string | X | √ |  | 
| default_value | string | X | √ |  | 
| name | string | X | √ |  | 
| azure_postgresql_servers_selefra_id | string | X | X | fk to azure_postgresql_servers.selefra_id | 
| value | string | X | √ |  | 
| data_type | string | X | √ |  | 
| allowed_values | string | X | √ |  | 
| source | string | X | √ |  | 
| id | string | √ | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


