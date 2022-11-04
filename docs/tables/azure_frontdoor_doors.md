# Table: azure_frontdoor_doors

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| load_balancing_settings | json | X | √ |  | 
| frontend_endpoints | json | X | √ |  | 
| id | string | √ | √ |  | 
| location | string | X | √ |  | 
| health_probe_settings | json | X | √ |  | 
| backend_pools | json | X | √ |  | 
| backend_pools_settings | json | X | √ |  | 
| enabled_state | string | X | √ |  | 
| tags | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| cname | string | X | √ |  | 
| frontdoor_id | string | X | √ |  | 
| rules_engines | json | X | √ |  | 
| routing_rules | json | X | √ |  | 
| type | string | X | √ |  | 
| resource_state | string | X | √ |  | 
| friendly_name | string | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


