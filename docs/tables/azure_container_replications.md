# Table: azure_container_replications

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| type | string | X | √ |  | 
| location | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| status | json | X | √ |  | 
| tags | json | X | √ |  | 
| azure_container_registries_selefra_id | string | X | X | fk to azure_container_registries.selefra_id | 
| subscription_id | string | X | √ |  | 
| container_registry_id | string | X | √ |  | 


