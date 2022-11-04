# Table: azure_network_security_groups

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subscription_id | string | X | √ |  | 
| id | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| security_rules | json | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| name | string | X | √ |  | 
| location | string | X | √ |  | 
| default_security_rules | json | X | √ |  | 
| subnets | json | X | √ |  | 
| resource_guid | string | X | √ |  | 
| type | string | X | √ |  | 
| network_interfaces | json | X | √ |  | 
| flow_logs | json | X | √ |  | 
| etag | string | X | √ |  | 
| tags | json | X | √ |  | 


