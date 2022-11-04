# Table: azure_mariadb_configurations

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subscription_id | string | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| azure_mariadb_servers_selefra_id | string | X | X | fk to azure_mariadb_servers.selefra_id | 
| allowed_values | string | X | √ |  | 
| source | string | X | √ |  | 
| id | string | √ | √ |  | 
| mariadb_server_id | string | X | √ |  | 
| value | string | X | √ |  | 
| description | string | X | √ |  | 
| default_value | string | X | √ |  | 
| data_type | string | X | √ |  | 
| name | string | X | √ |  | 


