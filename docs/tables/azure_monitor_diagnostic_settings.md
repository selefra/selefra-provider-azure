# Table: azure_monitor_diagnostic_settings

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| metrics | json | X | √ |  | 
| workspace_id | string | X | √ |  | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| subscription_id | string | X | √ |  | 
| monitor_resource_id | string | X | √ |  | 
| service_bus_rule_id | string | X | √ |  | 
| event_hub_authorization_rule_id | string | X | √ |  | 
| log_analytics_destination_type | string | X | √ |  | 
| resource_uri | string | X | √ |  | 
| storage_account_id | string | X | √ |  | 
| id | string | √ | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| event_hub_name | string | X | √ |  | 
| logs | json | X | √ |  | 
| azure_monitor_resources_selefra_id | string | X | X | fk to azure_monitor_resources.selefra_id | 


