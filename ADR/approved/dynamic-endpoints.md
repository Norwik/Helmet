### Dynamic Endpoints

Endpoints has to be defined in helmet config file. Anytime we want to add a new endpoint or change an endpoint, a restart will be needed. This draft explains how to make these configs dynamic.

- [ ] First the database schema need to be adjusted

```zsh
+ option:
---------
    - id
    - key
    - value

+ endpoint:
-----------
    - id
    - status
    - listen_path
    - name
    - upstreams
    - balancing
    - http_methods
    - authentication
    - rate_limit
    - circuit_breaker

+ auth_method:
--------------
    - id
    - name
    - description
    - type (key_authentication, basic_authentication, oauth_authentication, any_authentication)

+ endpoint_auth_method (Many to Many):
--------------------------------------
    - id
    - auth_method_id
    - endpoint_id

+ key_based_auth_data:
----------------------
	- id
	- name
	- api_key
	- meta
	- auth_method_id

+ basic_auth_data:
------------------
	- id
	- name
	- username
	- password
	- meta
	- auth_method_id

+ oauth_data:
-------------
	- id
	- name
	- client_id
	- client_secret
	- meta
	- auth_method_id

+ oauth_access_data:
--------------------
	- id
	- access_token
	- meta
	- expire_at
	- oauth_data_id
```

- [ ] Services and Endpoints has to be adjusted
