# Table: azure_cosmosdb_sql_databases

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cosmosdb_account_id | string | X | √ |  | 
| type | string | X | √ |  | 
| tags | json | X | √ |  | 
| azure_cosmosdb_accounts_selefra_id | string | X | X | fk to azure_cosmosdb_accounts.selefra_id | 
| name | string | X | √ |  | 
| location | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| subscription_id | string | X | √ |  | 
| resource | json | X | √ |  | 
| options | json | X | √ |  | 
| id | string | √ | √ |  | 


