# Table: azure_cdn_rules

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cdn_rule_set_id | string | X | √ |  | 
| match_processing_behavior | string | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| id | string | √ | √ |  | 
| subscription_id | string | X | √ |  | 
| actions | json | X | √ |  | 
| system_data | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| order | int | X | √ |  | 
| conditions | json | X | √ |  | 
| deployment_status | string | X | √ |  | 
| type | string | X | √ |  | 
| name | string | X | √ |  | 
| azure_cdn_rule_sets_selefra_id | string | X | X | fk to azure_cdn_rule_sets.selefra_id | 


