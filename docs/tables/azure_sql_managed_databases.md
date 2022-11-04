# Table: azure_sql_managed_databases

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| source_database_id | string | X | √ |  | 
| restorable_dropped_database_id | string | X | √ |  | 
| long_term_retention_backup_resource_id | string | X | √ |  | 
| name | string | X | √ |  | 
| earliest_restore_point | timestamp | X | √ |  | 
| recoverable_database_id | string | X | √ |  | 
| id | string | √ | √ |  | 
| type | string | X | √ |  | 
| create_mode | string | X | √ |  | 
| storage_container_sas_token | string | X | √ |  | 
| failover_group_id | string | X | √ |  | 
| last_backup_name | string | X | √ |  | 
| sql_managed_instance_id | string | X | √ |  | 
| restore_point_in_time | timestamp | X | √ |  | 
| default_secondary_location | string | X | √ |  | 
| catalog_collation | string | X | √ |  | 
| storage_container_uri | string | X | √ |  | 
| auto_complete_restore | bool | X | √ |  | 
| location | string | X | √ |  | 
| tags | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| collation | string | X | √ |  | 
| status | string | X | √ |  | 
| creation_date | timestamp | X | √ |  | 
| azure_sql_managed_instances_selefra_id | string | X | X | fk to azure_sql_managed_instances.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 


