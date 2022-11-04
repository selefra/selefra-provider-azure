# Table: azure_keyvault_vaults

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| properties_vault_uri | string | X | √ |  | 
| properties_enabled_for_deployment | bool | X | √ |  | 
| properties_soft_delete_retention_in_days | int | X | √ |  | 
| properties_enable_soft_delete | bool | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| location | string | X | √ |  | 
| tags | json | X | √ |  | 
| properties_tenant_id | string | X | √ |  | 
| properties_enabled_for_disk_encryption | bool | X | √ |  | 
| properties_enabled_for_template_deployment | bool | X | √ |  | 
| properties_create_mode | string | X | √ |  | 
| properties_enable_purge_protection | bool | X | √ |  | 
| properties_network_acls | json | X | √ |  | 
| properties_private_endpoint_connections | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| type | string | X | √ |  | 
| properties_sku | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| properties_access_policies | json | X | √ |  | 
| properties_enable_rbac_authorization | bool | X | √ |  | 


