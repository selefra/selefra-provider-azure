# Table: azure_redis_caches

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| replicas_per_master | int | X | √ |  | 
| shard_count | int | X | √ |  | 
| name | string | X | √ |  | 
| tenant_settings | json | X | √ |  | 
| zones | string_array | X | √ |  | 
| tags | json | X | √ |  | 
| public_network_access | string | X | √ |  | 
| id | string | √ | √ |  | 
| subnet_id | string | X | √ |  | 
| location | string | X | √ |  | 
| linked_servers | json | X | √ |  | 
| private_endpoint_connections | json | X | √ |  | 
| replicas_per_primary | int | X | √ |  | 
| access_keys | json | X | √ |  | 
| enable_non_ssl_port | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| provisioning_state | string | X | √ |  | 
| host_name | string | X | √ |  | 
| port | int | X | √ |  | 
| static_ip | string | X | √ |  | 
| redis_version | string | X | √ |  | 
| minimum_tls_version | string | X | √ |  | 
| subscription_id | string | X | √ |  | 
| ssl_port | int | X | √ |  | 
| instances | json | X | √ |  | 
| sku | json | X | √ |  | 
| redis_configuration | json | X | √ |  | 
| type | string | X | √ |  | 


