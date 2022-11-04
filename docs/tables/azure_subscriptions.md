# Table: azure_subscriptions

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| display_name | string | X | √ |  | 
| state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| managed_by_tenants | json | X | √ |  | 
| tags | json | X | √ |  | 
| id | string | √ | √ |  | 
| tenant_id | string | X | √ |  | 
| authorization_source | string | X | √ |  | 
| subscription_policies | json | X | √ |  | 


