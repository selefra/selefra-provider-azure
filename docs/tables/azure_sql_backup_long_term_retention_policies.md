# Table: azure_sql_backup_long_term_retention_policies

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subscription_id | string | X | √ |  | 
| sql_database_id | string | X | √ |  | 
| weekly_retention | string | X | √ |  | 
| week_of_year | int | X | √ |  | 
| id | string | √ | √ |  | 
| monthly_retention | string | X | √ |  | 
| yearly_retention | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| azure_sql_databases_selefra_id | string | X | X | fk to azure_sql_databases.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 


