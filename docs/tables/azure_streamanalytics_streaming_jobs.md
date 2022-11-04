# Table: azure_streamanalytics_streaming_jobs

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| job_state | string | X | √ |  | 
| tags | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| job_type | string | X | √ |  | 
| last_output_event_time | timestamp | X | √ |  | 
| data_locale | string | X | √ |  | 
| created_date | timestamp | X | √ |  | 
| content_storage_policy | string | X | √ |  | 
| cluster | json | X | √ |  | 
| id | string | √ | √ |  | 
| events_out_of_order_policy | string | X | √ |  | 
| inputs | json | X | √ |  | 
| job_storage_account | json | X | √ |  | 
| outputs | json | X | √ |  | 
| identity | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| job_id | string | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| output_error_policy | string | X | √ |  | 
| events_out_of_order_max_delay_in_seconds | int | X | √ |  | 
| compatibility_level | string | X | √ |  | 
| etag | string | X | √ |  | 
| name | string | X | √ |  | 
| output_start_time | timestamp | X | √ |  | 
| location | string | X | √ |  | 
| sku | json | X | √ |  | 
| output_start_mode | string | X | √ |  | 
| events_late_arrival_max_delay_in_seconds | int | X | √ |  | 
| transformation | json | X | √ |  | 
| functions | json | X | √ |  | 
| type | string | X | √ |  | 


