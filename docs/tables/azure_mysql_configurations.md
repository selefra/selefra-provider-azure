# Table: azure_mysql_configurations

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | √ | √ |  | 
| mysql_server_id | string | X | √ |  | 
| value | string | X | √ |  | 
| description | string | X | √ |  | 
| default_value | string | X | √ |  | 
| data_type | string | X | √ |  | 
| allowed_values | string | X | √ |  | 
| source | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| subscription_id | string | X | √ |  | 
| azure_mysql_servers_selefra_id | string | X | X | fk to azure_mysql_servers.selefra_id | 


