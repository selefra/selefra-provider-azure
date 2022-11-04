# Table: azure_web_publishing_profiles

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| user_pwd | string | X | √ |  | 
| azure_web_apps_selefra_id | string | X | X | fk to azure_web_apps.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| subscription_id | string | X | √ |  | 
| web_app_id | string | X | √ |  | 
| publish_url | string | X | √ |  | 
| user_name | string | X | √ |  | 


