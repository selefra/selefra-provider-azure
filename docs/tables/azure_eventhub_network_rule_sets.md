# Table: azure_eventhub_network_rule_sets

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| azure_eventhub_namespaces_selefra_id | string | X | X | fk to azure_eventhub_namespaces.selefra_id | 
| eventhub_namespace_id | string | X | √ |  | 
| default_action | string | X | √ |  | 
| type | string | X | √ |  | 
| ip_rules | json | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| subscription_id | string | X | √ |  | 
| trusted_service_access_enabled | bool | X | √ |  | 
| virtual_network_rules | json | X | √ |  | 


