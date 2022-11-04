# Table: azure_datalake_store_accounts

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_modified_time | timestamp | X | √ |  | 
| subscription_id | string | X | √ |  | 
| encryption_config | json | X | √ |  | 
| virtual_network_rules | json | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| name | string | X | √ |  | 
| tags | json | X | √ |  | 
| identity | json | X | √ |  | 
| firewall_allow_azure_ips | string | X | √ |  | 
| new_tier | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| endpoint | string | X | √ |  | 
| id | string | √ | √ |  | 
| type | string | X | √ |  | 
| encryption_provisioning_state | string | X | √ |  | 
| firewall_rules | json | X | √ |  | 
| trusted_id_providers | json | X | √ |  | 
| current_tier | string | X | √ |  | 
| account_id | string | X | √ |  | 
| state | string | X | √ |  | 
| location | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| default_group | string | X | √ |  | 
| encryption_state | string | X | √ |  | 
| firewall_state | string | X | √ |  | 
| trusted_id_provider_state | string | X | √ |  | 


