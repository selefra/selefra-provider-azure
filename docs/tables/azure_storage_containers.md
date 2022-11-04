# Table: azure_storage_containers

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| has_legal_hold | bool | X | √ |  | 
| has_immutability_policy | bool | X | √ |  | 
| etag | string | X | √ |  | 
| version | string | X | √ |  | 
| deleted_time | timestamp | X | √ |  | 
| last_modified_time | timestamp | X | √ |  | 
| metadata | json | X | √ |  | 
| immutability_policy | json | X | √ |  | 
| name | string | X | √ |  | 
| azure_storage_accounts_selefra_id | string | X | X | fk to azure_storage_accounts.selefra_id | 
| subscription_id | string | X | √ |  | 
| deleted | bool | X | √ |  | 
| default_encryption_scope | string | X | √ |  | 
| lease_status | string | X | √ |  | 
| legal_hold | json | X | √ |  | 
| storage_account_id | string | X | √ |  | 
| deny_encryption_scope_override | bool | X | √ |  | 
| lease_duration | string | X | √ |  | 
| id | string | √ | √ |  | 
| remaining_retention_days | int | X | √ |  | 
| public_access | string | X | √ |  | 
| lease_state | string | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


