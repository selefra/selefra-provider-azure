# Table: azure_network_route_tables

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| disable_bgp_route_propagation | bool | X | √ |  | 
| type | string | X | √ |  | 
| location | string | X | √ |  | 
| tags | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| routes | json | X | √ |  | 
| resource_guid | string | X | √ |  | 
| etag | string | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| subnets | json | X | √ |  | 
| provisioning_state | string | X | √ |  | 


