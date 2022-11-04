# Table: azure_postgresql_servers

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| user_visible_state | string | X | √ |  | 
| fully_qualified_domain_name | string | X | √ |  | 
| earliest_restore_date | timestamp | X | √ |  | 
| replication_role | string | X | √ |  | 
| location | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| ssl_enforcement | string | X | √ |  | 
| infrastructure_encryption | string | X | √ |  | 
| replica_capacity | int | X | √ |  | 
| tags | json | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| administrator_login | string | X | √ |  | 
| minimal_tls_version | string | X | √ |  | 
| version | string | X | √ |  | 
| byok_enforcement | string | X | √ |  | 
| storage_profile | json | X | √ |  | 
| master_server_id | string | X | √ |  | 
| identity | json | X | √ |  | 
| sku | json | X | √ |  | 
| private_endpoint_connections | json | X | √ |  | 
| type | string | X | √ |  | 
| subscription_id | string | X | √ |  | 
| public_network_access | string | X | √ |  | 


