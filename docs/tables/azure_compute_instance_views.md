# Table: azure_compute_instance_views

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| rdp_thumb_print | string | X | √ |  | 
| azure_compute_virtual_machines_selefra_id | string | X | X | fk to azure_compute_virtual_machines.selefra_id | 
| os_version | string | X | √ |  | 
| os_name | string | X | √ |  | 
| patch_status | json | X | √ |  | 
| compute_virtual_machine_id | string | X | √ |  | 
| computer_name | string | X | √ |  | 
| hyper_v_generation | string | X | √ |  | 
| maintenance_redeploy_status | json | X | √ |  | 
| vm_health | json | X | √ |  | 
| boot_diagnostics | json | X | √ |  | 
| assigned_host | string | X | √ |  | 
| statuses | json | X | √ |  | 
| platform_update_domain | int | X | √ |  | 
| platform_fault_domain | int | X | √ |  | 
| vm_agent | json | X | √ |  | 
| disks | json | X | √ |  | 
| extensions | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| subscription_id | string | X | √ |  | 


