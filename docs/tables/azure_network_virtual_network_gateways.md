# Table: azure_network_virtual_network_gateways

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| provisioning_state | string | X | √ |  | 
| ip_configurations | json | X | √ |  | 
| vpn_type | string | X | √ |  | 
| sku | json | X | √ |  | 
| network_virtual_network_id | string | X | √ |  | 
| enable_bgp | bool | X | √ |  | 
| inbound_dns_forwarding_endpoint | string | X | √ |  | 
| tags | json | X | √ |  | 
| azure_network_virtual_networks_selefra_id | string | X | X | fk to azure_network_virtual_networks.selefra_id | 
| gateway_default_site | json | X | √ |  | 
| location | string | X | √ |  | 
| etag | string | X | √ |  | 
| subscription_id | string | X | √ |  | 
| vpn_gateway_generation | string | X | √ |  | 
| bgp_settings | json | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| extended_location | json | X | √ |  | 
| gateway_type | string | X | √ |  | 
| enable_private_ip_address | bool | X | √ |  | 
| vpn_client_configuration | json | X | √ |  | 
| custom_routes | json | X | √ |  | 
| resource_guid | string | X | √ |  | 
| active_active | bool | X | √ |  | 
| enable_dns_forwarding | bool | X | √ |  | 
| v_net_extended_location_resource_id | string | X | √ |  | 
| id | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


