# Table: azure_network_public_ip_addresses

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| sku | json | X | √ |  | 
| ip_address | string | X | √ |  | 
| nat_gateway | json | X | √ |  | 
| etag | string | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| linked_public_ip_address | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| extended_location | json | X | √ |  | 
| public_ip_allocation_method | string | X | √ |  | 
| public_ip_address_version | string | X | √ |  | 
| ip_configuration | json | X | √ |  | 
| dns_settings | json | X | √ |  | 
| location | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| ip_tags | json | X | √ |  | 
| idle_timeout_in_minutes | int | X | √ |  | 
| resource_guid | string | X | √ |  | 
| service_public_ip_address | json | X | √ |  | 
| migration_phase | string | X | √ |  | 
| zones | string_array | X | √ |  | 
| ddos_settings | json | X | √ |  | 
| public_ip_prefix | json | X | √ |  | 
| type | string | X | √ |  | 


