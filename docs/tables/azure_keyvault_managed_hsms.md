# Table: azure_keyvault_managed_hsms

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| properties_enable_purge_protection | bool | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| subscription_id | string | X | √ |  | 
| properties_enable_soft_delete | bool | X | √ |  | 
| properties_provisioning_state | string | X | √ |  | 
| id | string | √ | √ |  | 
| type | string | X | √ |  | 
| tags | json | X | √ |  | 
| properties_hsm_uri | string | X | √ |  | 
| properties_soft_delete_retention_in_days | int | X | √ |  | 
| properties_status_message | string | X | √ |  | 
| properties_tenant_id | string | X | √ |  | 
| properties_initial_admin_object_ids | string_array | X | √ |  | 
| properties_create_mode | string | X | √ |  | 
| location | string | X | √ |  | 
| sku | json | X | √ |  | 


