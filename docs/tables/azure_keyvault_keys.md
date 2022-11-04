# Table: azure_keyvault_keys

## Primary Keys 

```
kid
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kid | string | √ | √ |  | 
| attributes | json | X | √ |  | 
| tags | json | X | √ |  | 
| managed | bool | X | √ |  | 
| azure_keyvault_vaults_selefra_id | string | X | X | fk to azure_keyvault_vaults.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| subscription_id | string | X | √ |  | 
| keyvault_vault_id | string | X | √ |  | 


