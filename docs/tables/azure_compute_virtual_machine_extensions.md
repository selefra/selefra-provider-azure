# Table: azure_compute_virtual_machine_extensions

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| compute_virtual_machine_id | string | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| name | string | X | √ |  | 
| location | string | X | √ |  | 
| force_update_tag | string | X | √ |  | 
| auto_upgrade_minor_version | bool | X | √ |  | 
| type | string | X | √ |  | 
| azure_compute_virtual_machines_selefra_id | string | X | X | fk to azure_compute_virtual_machines.selefra_id | 
| publisher | string | X | √ |  | 
| enable_automatic_upgrade | bool | X | √ |  | 
| instance_view | json | X | √ |  | 
| id | string | √ | √ |  | 
| subscription_id | string | X | √ |  | 
| type_handler_version | string | X | √ |  | 
| tags | json | X | √ |  | 


