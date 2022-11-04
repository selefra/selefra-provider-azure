# Table: azure_sql_databases

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| managed_by | string | X | √ |  | 
| source_database_deletion_date | timestamp | X | √ |  | 
| catalog_collation | string | X | √ |  | 
| earliest_restore_date | timestamp | X | √ |  | 
| tags | json | X | √ |  | 
| sql_server_id | string | X | √ |  | 
| sample_name | string | X | √ |  | 
| source_database_id | string | X | √ |  | 
| max_log_size_bytes | int | X | √ |  | 
| current_sku | json | X | √ |  | 
| max_size_bytes | int | X | √ |  | 
| license_type | string | X | √ |  | 
| resumed_date | timestamp | X | √ |  | 
| id | string | √ | √ |  | 
| subscription_id | string | X | √ |  | 
| kind | string | X | √ |  | 
| elastic_pool_id | string | X | √ |  | 
| status | string | X | √ |  | 
| database_id | string | X | √ |  | 
| current_service_objective_name | string | X | √ |  | 
| requested_service_objective_name | string | X | √ |  | 
| failover_group_id | string | X | √ |  | 
| restorable_dropped_database_id | string | X | √ |  | 
| maintenance_configuration_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| sku | json | X | √ |  | 
| secondary_type | string | X | √ |  | 
| paused_date | timestamp | X | √ |  | 
| location | string | X | √ |  | 
| type | string | X | √ |  | 
| collation | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| restore_point_in_time | timestamp | X | √ |  | 
| recovery_services_recovery_point_id | string | X | √ |  | 
| zone_redundant | bool | X | √ |  | 
| read_scale | string | X | √ |  | 
| high_availability_replica_count | int | X | √ |  | 
| auto_pause_delay | int | X | √ |  | 
| min_capacity | float | X | √ |  | 
| recoverable_database_id | string | X | √ |  | 
| create_mode | string | X | √ |  | 
| default_secondary_location | string | X | √ |  | 
| long_term_retention_backup_resource_id | string | X | √ |  | 
| storage_account_type | string | X | √ |  | 
| name | string | X | √ |  | 
| azure_sql_servers_selefra_id | string | X | X | fk to azure_sql_servers.selefra_id | 


