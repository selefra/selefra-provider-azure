# Table: azure_cdn_custom_domains

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subscription_id | string | X | √ |  | 
| host_name | string | X | √ |  | 
| validation_data | string | X | √ |  | 
| name | string | X | √ |  | 
| azure_cdn_endpoints_selefra_id | string | X | X | fk to azure_cdn_endpoints.selefra_id | 
| resource_state | string | X | √ |  | 
| id | string | √ | √ |  | 
| custom_https_provisioning_substate | string | X | √ |  | 
| cdn_endpoint_id | string | X | √ |  | 
| custom_https_provisioning_state | string | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| type | string | X | √ |  | 
| system_data | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


