# Table: azure_network_flow_logs

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| network_watcher_id | string | X | √ |  | 
| target_resource_guid | string | X | √ |  | 
| enabled | bool | X | √ |  | 
| retention_policy | json | X | √ |  | 
| type | string | X | √ |  | 
| location | string | X | √ |  | 
| tags | json | X | √ |  | 
| storage_id | string | X | √ |  | 
| format | json | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| azure_network_watchers_selefra_id | string | X | X | fk to azure_network_watchers.selefra_id | 
| subscription_id | string | X | √ |  | 
| target_resource_id | string | X | √ |  | 
| etag | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| flow_analytics_configuration | json | X | √ |  | 
| provisioning_state | string | X | √ |  | 


