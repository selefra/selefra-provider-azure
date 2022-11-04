# Table: azure_resources_policy_assignments

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| type | string | X | √ |  | 
| sku | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| display_name | string | X | √ |  | 
| not_scopes | string_array | X | √ |  | 
| enforcement_mode | string | X | √ |  | 
| id | string | √ | √ |  | 
| identity | json | X | √ |  | 
| policy_definition_id | string | X | √ |  | 
| scope | string | X | √ |  | 
| location | string | X | √ |  | 
| parameters | json | X | √ |  | 
| name | string | X | √ |  | 


