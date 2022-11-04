# Table: azure_mariadb_servers

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| location | string | X | √ |  | 
| sku | json | X | √ |  | 
| administrator_login | string | X | √ |  | 
| earliest_restore_date | timestamp | X | √ |  | 
| storage_profile | json | X | √ |  | 
| master_server_id | string | X | √ |  | 
| private_endpoint_connections | json | X | √ |  | 
| tags | json | X | √ |  | 
| id | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| version | string | X | √ |  | 
| user_visible_state | string | X | √ |  | 
| type | string | X | √ |  | 
| ssl_enforcement | string | X | √ |  | 
| fully_qualified_domain_name | string | X | √ |  | 
| name | string | X | √ |  | 
| subscription_id | string | X | √ |  | 
| replication_role | string | X | √ |  | 
| replica_capacity | int | X | √ |  | 
| public_network_access | string | X | √ |  | 


