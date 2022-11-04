# Table: azure_keyvault_secrets

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| keyvault_vault_id | string | X | √ |  | 
| attributes | json | X | √ |  | 
| managed | bool | X | √ |  | 
| azure_keyvault_vaults_selefra_id | string | X | X | fk to azure_keyvault_vaults.selefra_id | 
| subscription_id | string | X | √ |  | 
| tags | json | X | √ |  | 
| content_type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | √ | √ |  | 


