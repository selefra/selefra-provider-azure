# Table: azure_sql_managed_instance_encryption_protectors

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| uri | string | X | √ |  | 
| subscription_id | string | X | √ |  | 
| sql_managed_instance_id | string | X | √ |  | 
| kind | string | X | √ |  | 
| server_key_name | string | X | √ |  | 
| server_key_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| thumbprint | string | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| azure_sql_managed_instances_selefra_id | string | X | X | fk to azure_sql_managed_instances.selefra_id | 


