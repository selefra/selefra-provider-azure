# Table: azure_mysql_servers

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| replica_capacity | int | X | √ |  | 
| private_endpoint_connections | json | X | √ |  | 
| sku | json | X | √ |  | 
| tags | json | X | √ |  | 
| location | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| subscription_id | string | X | √ |  | 
| administrator_login | string | X | √ |  | 
| byok_enforcement | string | X | √ |  | 
| user_visible_state | string | X | √ |  | 
| fully_qualified_domain_name | string | X | √ |  | 
| replication_role | string | X | √ |  | 
| id | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| identity | json | X | √ |  | 
| version | string | X | √ |  | 
| ssl_enforcement | string | X | √ |  | 
| minimal_tls_version | string | X | √ |  | 
| infrastructure_encryption | string | X | √ |  | 
| earliest_restore_date | timestamp | X | √ |  | 
| storage_profile | json | X | √ |  | 
| master_server_id | string | X | √ |  | 
| public_network_access | string | X | √ |  | 


