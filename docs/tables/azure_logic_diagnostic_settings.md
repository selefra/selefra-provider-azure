# Table: azure_logic_diagnostic_settings

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subscription_id | string | X | √ |  | 
| logic_workflow_id | string | X | √ |  | 
| metrics | json | X | √ |  | 
| storage_account_id | string | X | √ |  | 
| event_hub_name | string | X | √ |  | 
| workspace_id | string | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| service_bus_rule_id | string | X | √ |  | 
| id | string | √ | √ |  | 
| type | string | X | √ |  | 
| event_hub_authorization_rule_id | string | X | √ |  | 
| logs | json | X | √ |  | 
| log_analytics_destination_type | string | X | √ |  | 
| azure_logic_workflows_selefra_id | string | X | X | fk to azure_logic_workflows.selefra_id | 


