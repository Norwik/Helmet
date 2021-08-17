#### Dynamic Endpoints

Endpoints has to be defined in helmet config file. Anytime we want to add a new endpoint or change an endpoint, a restart will be needed. This draft explains how to make these configs dynamic.

- [ ] First the database schema need to be adjusted

```zsh
+ option:
    - id
    - key
    - value

+ endpoint:
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
    - id
    - name
    - description
    - type (key_authentication, basic_authentication, oauth_authentication, any_authentication)

+ endpoint_auth_method (Many to Many):
    - id
    - auth_method_id
    - endpoint_id
```

- [ ] Services and Endpoints has to be adjusted
