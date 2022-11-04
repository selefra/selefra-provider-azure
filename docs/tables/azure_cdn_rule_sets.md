# Table: azure_cdn_rule_sets

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| type | string | X | √ |  | 
| system_data | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| subscription_id | string | X | √ |  | 
| deployment_status | string | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| cdn_profile_id | string | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| azure_cdn_profiles_selefra_id | string | X | X | fk to azure_cdn_profiles.selefra_id | 


