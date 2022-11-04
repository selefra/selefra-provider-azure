# Table: azure_sql_managed_instances

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| provisioning_state | string | X | √ |  | 
| administrator_login_password | string | X | √ |  | 
| v_cores | int | X | √ |  | 
| restore_point_in_time | timestamp | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| collation | string | X | √ |  | 
| dns_zone_partner | string | X | √ |  | 
| private_endpoint_connections | json | X | √ |  | 
| zone_redundant | bool | X | √ |  | 
| type | string | X | √ |  | 
| subnet_id | string | X | √ |  | 
| license_type | string | X | √ |  | 
| dns_zone | string | X | √ |  | 
| public_data_endpoint_enabled | bool | X | √ |  | 
| tags | json | X | √ |  | 
| administrator_login | string | X | √ |  | 
| instance_pool_id | string | X | √ |  | 
| minimal_tls_version | string | X | √ |  | 
| subscription_id | string | X | √ |  | 
| state | string | X | √ |  | 
| timezone_id | string | X | √ |  | 
| location | string | X | √ |  | 
| identity | json | X | √ |  | 
| sku | json | X | √ |  | 
| storage_account_type | string | X | √ |  | 
| managed_instance_create_mode | string | X | √ |  | 
| fully_qualified_domain_name | string | X | √ |  | 
| storage_size_in_gb | int | X | √ |  | 
| source_managed_instance_id | string | X | √ |  | 
| proxy_override | string | X | √ |  | 
| maintenance_configuration_id | string | X | √ |  | 


