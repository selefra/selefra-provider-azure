# Table: azure_monitor_log_profiles

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| etag | string | X | √ |  | 
| locations | string_array | X | √ |  | 
| retention_policy | json | X | √ |  | 
| name | string | X | √ |  | 
| location | string | X | √ |  | 
| tags | json | X | √ |  | 
| kind | string | X | √ |  | 
| categories | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| storage_account_id | string | X | √ |  | 
| id | string | √ | √ |  | 
| subscription_id | string | X | √ |  | 
| service_bus_rule_id | string | X | √ |  | 
| type | string | X | √ |  | 


