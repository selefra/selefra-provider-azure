# Table: azure_container_registries

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| sku | json | X | √ |  | 
| id | string | √ | √ |  | 
| storage_account | json | X | √ |  | 
| location | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| creation_date | timestamp | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| policies | json | X | √ |  | 
| tags | json | X | √ |  | 
| type | string | X | √ |  | 
| subscription_id | string | X | √ |  | 
| login_server | string | X | √ |  | 
| status | json | X | √ |  | 
| admin_user_enabled | bool | X | √ |  | 
| network_rule_set | json | X | √ |  | 
| name | string | X | √ |  | 


