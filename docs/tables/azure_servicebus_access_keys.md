# Table: azure_servicebus_access_keys

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| primary_connection_string | string | X | √ |  | 
| secondary_connection_string | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| subscription_id | string | X | √ |  | 
| servicebus_authorization_rule_id | string | X | √ |  | 
| primary_key | string | X | √ |  | 
| secondary_key | string | X | √ |  | 
| key_name | string | X | √ |  | 
| azure_servicebus_authorization_rules_selefra_id | string | X | X | fk to azure_servicebus_authorization_rules.selefra_id | 
| alias_primary_connection_string | string | X | √ |  | 
| alias_secondary_connection_string | string | X | √ |  | 


