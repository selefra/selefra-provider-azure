# Table: azure_cdn_security_policies

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cdn_profile_id | string | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| system_data | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| deployment_status | string | X | √ |  | 
| azure_cdn_profiles_selefra_id | string | X | X | fk to azure_cdn_profiles.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| provisioning_state | string | X | √ |  | 


