# Table: azure_web_site_auth_settings

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| subscription_id | string | X | √ |  | 
| token_refresh_extension_hours | float | X | √ |  | 
| microsoft_account_client_id | string | X | √ |  | 
| microsoft_account_o_auth_scopes | string_array | X | √ |  | 
| token_store_enabled | bool | X | √ |  | 
| validate_issuer | bool | X | √ |  | 
| git_hub_client_id | string | X | √ |  | 
| web_app_id | string | X | √ |  | 
| allowed_external_redirect_urls | string_array | X | √ |  | 
| google_client_secret_setting_name | string | X | √ |  | 
| allowed_audiences | string_array | X | √ |  | 
| google_client_secret | string | X | √ |  | 
| facebook_o_auth_scopes | string_array | X | √ |  | 
| git_hub_client_secret_setting_name | string | X | √ |  | 
| twitter_consumer_key | string | X | √ |  | 
| client_secret_setting_name | string | X | √ |  | 
| client_secret_certificate_thumbprint | string | X | √ |  | 
| issuer | string | X | √ |  | 
| azure_web_apps_selefra_id | string | X | X | fk to azure_web_apps.selefra_id | 
| additional_login_params | string_array | X | √ |  | 
| google_o_auth_scopes | string_array | X | √ |  | 
| git_hub_o_auth_scopes | string_array | X | √ |  | 
| auth_file_path | string | X | √ |  | 
| config_version | string | X | √ |  | 
| unauthenticated_client_action | string | X | √ |  | 
| default_provider | string | X | √ |  | 
| client_id | string | X | √ |  | 
| name | string | X | √ |  | 
| kind | string | X | √ |  | 
| client_secret | string | X | √ |  | 
| is_auth_from_file | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| facebook_app_id | string | X | √ |  | 
| twitter_consumer_secret | string | X | √ |  | 
| microsoft_account_client_secret | string | X | √ |  | 
| type | string | X | √ |  | 
| runtime_version | string | X | √ |  | 
| aad_claims_authorization | string | X | √ |  | 
| google_client_id | string | X | √ |  | 
| git_hub_client_secret | string | X | √ |  | 
| twitter_consumer_secret_setting_name | string | X | √ |  | 
| microsoft_account_client_secret_setting_name | string | X | √ |  | 
| id | string | √ | √ |  | 
| enabled | bool | X | √ |  | 
| facebook_app_secret | string | X | √ |  | 
| facebook_app_secret_setting_name | string | X | √ |  | 


