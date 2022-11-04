# Table: azure_network_virtual_networks

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| enable_vm_protection | bool | X | √ |  | 
| ddos_protection_plan | json | X | √ |  | 
| ip_allocations | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| enable_ddos_protection | bool | X | √ |  | 
| type | string | X | √ |  | 
| location | string | X | √ |  | 
| address_space | json | X | √ |  | 
| subnets | json | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| bgp_communities | json | X | √ |  | 
| name | string | X | √ |  | 
| tags | json | X | √ |  | 
| extended_location | json | X | √ |  | 
| dhcp_options | json | X | √ |  | 
| virtual_network_peerings | json | X | √ |  | 
| resource_guid | string | X | √ |  | 
| etag | string | X | √ |  | 
| id | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


