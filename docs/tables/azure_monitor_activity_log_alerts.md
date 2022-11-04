# Table: azure_monitor_activity_log_alerts

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| scopes | string_array | X | √ |  | 
| enabled | bool | X | √ |  | 
| id | string | √ | √ |  | 
| description | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| location | string | X | √ |  | 
| subscription_id | string | X | √ |  | 
| condition | json | X | √ |  | 
| actions | json | X | √ |  | 


