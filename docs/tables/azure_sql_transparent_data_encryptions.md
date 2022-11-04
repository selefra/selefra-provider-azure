# Table: azure_sql_transparent_data_encryptions

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subscription_id | string | X | √ |  | 
| location | string | X | √ |  | 
| id | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| sql_database_id | string | X | √ |  | 
| status | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| azure_sql_databases_selefra_id | string | X | X | fk to azure_sql_databases.selefra_id | 


