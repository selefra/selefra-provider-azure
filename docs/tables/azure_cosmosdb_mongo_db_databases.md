# Table: azure_cosmosdb_mongo_db_databases

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| azure_cosmosdb_accounts_selefra_id | string | X | X | fk to azure_cosmosdb_accounts.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| subscription_id | string | X | √ |  | 
| options | json | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| location | string | X | √ |  | 
| cosmosdb_account_id | string | X | √ |  | 
| resource | json | X | √ |  | 
| type | string | X | √ |  | 
| tags | json | X | √ |  | 


