# Table: azure_web_vnet_connections

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| routes | json | X | √ |  | 
| kind | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| web_app_id | string | X | √ |  | 
| vnet_resource_id | string | X | √ |  | 
| cert_thumbprint | string | X | √ |  | 
| azure_web_apps_selefra_id | string | X | X | fk to azure_web_apps.selefra_id | 
| subscription_id | string | X | √ |  | 
| dns_servers | string | X | √ |  | 
| id | string | √ | √ |  | 
| type | string | X | √ |  | 
| cert_blob | string | X | √ |  | 
| resync_required | bool | X | √ |  | 
| is_swift | bool | X | √ |  | 
| name | string | X | √ |  | 


