# Table: azure_servicebus_authorization_rules

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subscription_id | string | X | √ |  | 
| servicebus_topic_id | string | X | √ |  | 
| system_data | json | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| azure_servicebus_topics_selefra_id | string | X | X | fk to azure_servicebus_topics.selefra_id | 
| rights | string_array | X | √ |  | 
| type | string | X | √ |  | 


