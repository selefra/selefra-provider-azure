# Table: azure_security_jit_network_access_policies

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| kind | string | X | √ |  | 
| requests | json | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| subscription_id | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| id | string | √ | √ |  | 
| location | string | X | √ |  | 
| virtual_machines | json | X | √ |  | 


