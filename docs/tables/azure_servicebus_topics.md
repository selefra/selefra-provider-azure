# Table: azure_servicebus_topics

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| azure_servicebus_namespaces_selefra_id | string | X | X | fk to azure_servicebus_namespaces.selefra_id | 
| subscription_id | string | X | √ |  | 
| requires_duplicate_detection | bool | X | √ |  | 
| enable_partitioning | bool | X | √ |  | 
| enable_express | bool | X | √ |  | 
| duplicate_detection_history_time_window | string | X | √ |  | 
| auto_delete_on_idle | string | X | √ |  | 
| subscription_count | int | X | √ |  | 
| count_details | json | X | √ |  | 
| max_size_in_megabytes | int | X | √ |  | 
| max_message_size_in_kilobytes | int | X | √ |  | 
| system_data | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| servicebus_namespace_id | string | X | √ |  | 
| size_in_bytes | int | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| support_ordering | bool | X | √ |  | 
| status | string | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| accessed_at | timestamp | X | √ |  | 
| default_message_time_to_live | string | X | √ |  | 
| enable_batched_operations | bool | X | √ |  | 


