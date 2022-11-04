# Table: azure_network_interfaces

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| migration_phase | string | X | √ |  | 
| etag | string | X | √ |  | 
| location | string | X | √ |  | 
| tags | json | X | √ |  | 
| ip_configurations | json | X | √ |  | 
| enable_ip_forwarding | bool | X | √ |  | 
| hosted_workloads | string_array | X | √ |  | 
| nic_type | string | X | √ |  | 
| primary | bool | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| subscription_id | string | X | √ |  | 
| network_security_group | json | X | √ |  | 
| private_endpoint | json | X | √ |  | 
| dns_settings | json | X | √ |  | 
| private_link_service | json | X | √ |  | 
| type | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| extended_location | json | X | √ |  | 
| tap_configurations | json | X | √ |  | 
| resource_guid | string | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| virtual_machine | json | X | √ |  | 
| mac_address | string | X | √ |  | 
| enable_accelerated_networking | bool | X | √ |  | 
| dscp_configuration | json | X | √ |  | 


