# Table: azure_monitor_activity_logs

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| operation_name | json | X | √ |  | 
| properties | json | X | √ |  | 
| status | json | X | √ |  | 
| resource_type | json | X | √ |  | 
| sub_status | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| subscription_id | string | X | √ |  | 
| authorization | json | X | √ |  | 
| id | string | √ | √ |  | 
| event_data_id | string | X | √ |  | 
| correlation_id | string | X | √ |  | 
| submission_timestamp | timestamp | X | √ |  | 
| tenant_id | string | X | √ |  | 
| description | string | X | √ |  | 
| category | json | X | √ |  | 
| level | string | X | √ |  | 
| resource_group_name | string | X | √ |  | 
| resource_provider_name | json | X | √ |  | 
| operation_id | string | X | √ |  | 
| event_timestamp | timestamp | X | √ |  | 
| claims | json | X | √ |  | 
| caller | string | X | √ |  | 
| event_name | json | X | √ |  | 
| http_request | json | X | √ |  | 
| resource_id | string | X | √ |  | 


