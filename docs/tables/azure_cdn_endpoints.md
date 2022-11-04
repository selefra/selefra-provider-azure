# Table: azure_cdn_endpoints

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| provisioning_state | string | X | √ |  | 
| origin_path | string | X | √ |  | 
| is_compression_enabled | bool | X | √ |  | 
| optimization_type | string | X | √ |  | 
| probe_path | string | X | √ |  | 
| tags | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| cdn_profile_id | string | X | √ |  | 
| origin_groups | json | X | √ |  | 
| is_https_allowed | bool | X | √ |  | 
| web_application_firewall_policy_link | json | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| host_name | string | X | √ |  | 
| resource_state | string | X | √ |  | 
| origin_host_header | string | X | √ |  | 
| type | string | X | √ |  | 
| azure_cdn_profiles_selefra_id | string | X | X | fk to azure_cdn_profiles.selefra_id | 
| origins | json | X | √ |  | 
| geo_filters | json | X | √ |  | 
| system_data | json | X | √ |  | 
| content_types_to_compress | string_array | X | √ |  | 
| default_origin_group | json | X | √ |  | 
| delivery_policy | json | X | √ |  | 
| query_string_caching_behavior | string | X | √ |  | 
| url_signing_keys | json | X | √ |  | 
| location | string | X | √ |  | 
| is_http_allowed | bool | X | √ |  | 


