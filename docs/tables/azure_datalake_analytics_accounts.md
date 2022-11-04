# Table: azure_datalake_analytics_accounts

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| last_modified_time | timestamp | X | √ |  | 
| max_job_count | int | X | √ |  | 
| system_max_job_count | int | X | √ |  | 
| account_id | string | X | √ |  | 
| provisioning_state | string | X | √ |  | 
| new_tier | string | X | √ |  | 
| min_priority_per_job | int | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| storage_accounts | json | X | √ |  | 
| firewall_allow_azure_ips | string | X | √ |  | 
| system_max_degree_of_parallelism | int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| state | string | X | √ |  | 
| creation_time | timestamp | X | √ |  | 
| firewall_rules | json | X | √ |  | 
| query_store_retention | int | X | √ |  | 
| endpoint | string | X | √ |  | 
| location | string | X | √ |  | 
| data_lake_store_accounts | json | X | √ |  | 
| compute_policies | json | X | √ |  | 
| max_degree_of_parallelism_per_job | int | X | √ |  | 
| type | string | X | √ |  | 
| tags | json | X | √ |  | 
| subscription_id | string | X | √ |  | 
| default_data_lake_store_account | string | X | √ |  | 
| max_degree_of_parallelism | int | X | √ |  | 
| firewall_state | string | X | √ |  | 
| current_tier | string | X | √ |  | 


