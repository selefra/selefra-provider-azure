# Table: azure_eventhub_namespaces

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| sku | json | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| service_bus_endpoint | string | X | √ |  | 
| zone_redundant | bool | X | √ |  | 
| is_auto_inflate_enabled | bool | X | √ |  | 
| maximum_throughput_units | int | X | √ |  | 
| encryption | json | X | √ |  | 
| tags | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| identity | json | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| metric_id | string | X | √ |  | 
| id | string | √ | √ |  | 
| updated_at | timestamp | X | √ |  | 
| kafka_enabled | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| cluster_arm_id | string | X | √ |  | 
| location | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 


