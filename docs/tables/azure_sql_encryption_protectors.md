# Table: azure_sql_encryption_protectors

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| uri | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| azure_sql_servers_selefra_id | string | X | X | fk to azure_sql_servers.selefra_id | 
| subscription_id | string | X | √ |  | 
| sql_server_id | string | X | √ |  | 
| server_key_name | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| location | string | X | √ |  | 
| subregion | string | X | √ |  | 
| thumbprint | string | X | √ |  | 
| id | string | √ | √ |  | 
| kind | string | X | √ |  | 
| server_key_type | string | X | √ |  | 


