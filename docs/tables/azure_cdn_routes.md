# Table: azure_cdn_routes

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| origin_group | json | X | √ |  | 
| deployment_status | string | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| custom_domains | json | X | √ |  | 
| origin_path | string | X | √ |  | 
| rule_sets | json | X | √ |  | 
| patterns_to_match | string_array | X | √ |  | 
| forwarding_protocol | string | X | √ |  | 
| system_data | json | X | √ |  | 
| supported_protocols | string_array | X | √ |  | 
| link_to_default_domain | string | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| id | string | √ | √ |  | 
| type | string | X | √ |  | 
| subscription_id | string | X | √ |  | 
| cdn_endpoint_id | string | X | √ |  | 
| query_string_caching_behavior | string | X | √ |  | 
| https_redirect | string | X | √ |  | 
| enabled_state | string | X | √ |  | 
| azure_cdn_endpoints_selefra_id | string | X | X | fk to azure_cdn_endpoints.selefra_id | 


