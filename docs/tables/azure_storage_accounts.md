# Table: azure_storage_accounts

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kind | string | X | √ |  | 
| primary_location | string | X | √ |  | 
| status_of_primary | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| secondary_endpoints | json | X | √ |  | 
| identity | json | X | √ |  | 
| extended_location | json | X | √ |  | 
| access_tier | string | X | √ |  | 
| is_hns_enabled | bool | X | √ |  | 
| geo_replication_stats | json | X | √ |  | 
| failover_in_progress | bool | X | √ |  | 
| private_endpoint_connections | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| subscription_id | string | X | √ |  | 
| sku | json | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| routing_preference | json | X | √ |  | 
| network_acls | json | X | √ |  | 
| blob_restore_status | json | X | √ |  | 
| minimum_tls_version | string | X | √ |  | 
| location | string | X | √ |  | 
| name | string | X | √ |  | 
| queue_logging_settings | json | X | √ |  | 
| is_nfs_v3_enabled | bool | X | √ |  | 
| last_geo_failover_time | timestamp | X | √ |  | 
| status_of_secondary | string | X | √ |  | 
| secondary_location | string | X | √ |  | 
| azure_files_identity_based_authentication | json | X | √ |  | 
| supports_https_traffic_only | bool | X | √ |  | 
| large_file_shares_state | string | X | √ |  | 
| allow_shared_key_access | bool | X | √ |  | 
| tags | json | X | √ |  | 
| id | string | √ | √ |  | 
| custom_domain | json | X | √ |  | 
| encryption | json | X | √ |  | 
| allow_blob_public_access | bool | X | √ |  | 
| type | string | X | √ |  | 
| primary_endpoints | json | X | √ |  | 
| blob_logging_settings | json | X | √ |  | 


