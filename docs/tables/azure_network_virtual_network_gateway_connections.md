# Table: azure_network_virtual_network_gateway_connections

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| subscription_id | string | X | √ |  | 
| routing_weight | int | X | √ |  | 
| traffic_selector_policies | json | X | √ |  | 
| connection_type | string | X | √ |  | 
| tunnel_connection_status | json | X | √ |  | 
| enable_bgp | bool | X | √ |  | 
| virtual_network_gateway2 | json | X | √ |  | 
| shared_key | string | X | √ |  | 
| type | string | X | √ |  | 
| authorization_key | string | X | √ |  | 
| local_network_gateway2 | json | X | √ |  | 
| connection_protocol | string | X | √ |  | 
| name | string | X | √ |  | 
| tags | json | X | √ |  | 
| azure_network_virtual_network_gateways_selefra_id | string | X | X | fk to azure_network_virtual_network_gateways.selefra_id | 
| network_virtual_network_gateway_id | string | X | √ |  | 
| connection_mode | string | X | √ |  | 
| resource_guid | string | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| etag | string | X | √ |  | 
| id | string | √ | √ |  | 
| egress_bytes_transferred | int | X | √ |  | 
| peer | json | X | √ |  | 
| ipsec_policies | json | X | √ |  | 
| location | string | X | √ |  | 
| virtual_network_gateway1 | json | X | √ |  | 
| connection_status | string | X | √ |  | 
| use_policy_based_traffic_selectors | bool | X | √ |  | 
| ingress_bytes_transferred | int | X | √ |  | 
| express_route_gateway_bypass | bool | X | √ |  | 


