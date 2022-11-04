# Table: azure_authorization_role_definitions

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| assignable_scopes | string_array | X | √ |  | 
| role_type | string | X | √ |  | 
| subscription_id | string | X | √ |  | 
| role_name | string | X | √ |  | 
| description | string | X | √ |  | 
| permissions | json | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


