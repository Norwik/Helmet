<p align="center">
    <img src="/assets/logo.png?v=1.0.10" width="200" />
    <h3 align="center">Helmet</h3>
    <p align="center">A Lightweight Cloud Native API Gateway.</p>
    <p align="center">
        <a href="https://github.com/spacewalkio/Helmet/actions/workflows/build.yml">
            <img src="https://github.com/spacewalkio/Helmet/actions/workflows/build.yml/badge.svg">
        </a>
        <a href="https://github.com/spacewalkio/Helmet/actions">
            <img src="https://github.com/spacewalkio/Helmet/workflows/Release/badge.svg">
        </a>
        <a href="https://github.com/spacewalkio/Helmet/releases">
            <img src="https://img.shields.io/badge/Version-1.0.10-9B59B6.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/spacewalkio/Helmet">
            <img src="https://goreportcard.com/badge/github.com/spacewalkio/Helmet?v=1.0.10">
        </a>
        <a href="https://godoc.org/github.com/spacewalkio/helmet">
            <img src="https://godoc.org/github.com/spacewalkio/helmet?status.svg">
        </a>
        <a href="https://github.com/spacewalkio/Helmet/blob/master/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-E74C3C.svg">
        </a>
    </p>
</p>
<br/>
<p align="center">
    <img src="/assets/chart.png?v=1.0.10" width="80%" />
</p>

Helmet is Cloud Native API Gateway that control who accesses your API whether from customer or other internal services. It also collect metrics about service calls count, latency, success rate and much more!

Here is some of the key features:

- Manage Service to Service Authentication, Authorization and Communication.
- Manage End User to Service Authentication, Authorization and Communication.
- Basic, API Key Based and OAuth2 Authentication Support.
- Multiple Backends Support with Load Balancing, Health Checks.
- Runtime Metrics for both Helmet and Backend Services exposed for Prometheus.
- CORS Support.
- HTTP/2 support.
- Rate Limiting Support.
- Circuit Breaker Support.
- Caching Layer to make it even more faster.
- Lightweight, Easy to Deploy and Operate.


## Documentation

#### Linux Deployment

Download [the latest helmet binary](https://github.com/spacewalkio/Helmet/releases). Make it executable from everywhere.

```zsh
$ export LATEST_VERSION=$(curl --silent "https://api.github.com/repos/spacewalkio/Helmet/releases/latest" \
    | jq '.tag_name' \
    | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

$ curl -sL https://github.com/spacewalkio/Helmet/releases/download/v{$LATEST_VERSION}/helmet_{$LATEST_VERSION}_Linux_x86_64.tar.gz \
    | tar xz
```

Then install `MySQL` and `Redis` on the server or a separate one.

Create the configs file `config.yml` from `config.dist.yml`. Something like the following:

```yaml
# App configs
app:
    # App name
    name: ${HELMET_APP_NAME:-helmet}
    # Env mode (dev or prod)
    mode: ${HELMET_APP_MODE:-dev}
    # HTTP port
    port: ${HELMET_API_PORT:-8000}
    # Hostname
    hostname: ${HELMET_API_HOSTNAME:-127.0.0.1}
    # TLS configs
    tls:
        status: ${HELMET_API_TLS_STATUS:-off}
        crt_path: ${HELMET_API_TLS_PEMPATH:-cert/server.crt}
        key_path: ${HELMET_API_TLS_KEYPATH:-cert/server.key}

    # Global timeout
    timeout: ${HELMET_API_TIMEOUT:-50}

    # API Configs
    api:
        key: ${HELMET_API_KEY:-6c68b836-6f8e-465e-b59f-89c1db53afca}

    # CORS status
    cors:
        status: ${HELMET_CORS_STATUS:-off}

    # Application Database
    database:
        # Database driver (sqlite3, mysql)
        driver: ${HELMET_DATABASE_DRIVER:-sqlite3}
        # Database Host
        host: ${HELMET_DATABASE_MYSQL_HOST:-localhost}
        # Database Port
        port: ${HELMET_DATABASE_MYSQL_PORT:-3306}
        # Database Name
        name: ${HELMET_DATABASE_MYSQL_DATABASE:-helmet.db}
        # Database Username
        username: ${HELMET_DATABASE_MYSQL_USERNAME:-root}
        # Database Password
        password: ${HELMET_DATABASE_MYSQL_PASSWORD:-root}

    # Key Store Configs
    key_store:
        # Cache Driver
        driver: ${HELMET_CACHE_DRIVER:-redis}
        # Redis Driver Configs
        redis:
            # Redis Address
            address: ${HELMET_CACHE_REDIS_ADDR:-localhost:6379}
            # Redis Password
            password: ${HELMET_CACHE_REDIS_PASSWORD:-}
            # Redis Database
            database: ${HELMET_CACHE_REDIS_DB:-0}

    # Endpoint Configs
    endpoint:
        # Orders Internal Service
        - name: order_service
          active: true
          proxy:
            listen_path: "/orders/v2/*"
            upstreams:
                balancing: roundrobin
                targets:
                    - target: https://httpbin.org/anything/orderService1/v2
                    - target: https://httpbin.org/anything/orderService2/v2
                    - target: https://httpbin.org/anything/orderService3/v2
                    - target: https://httpbin.org/anything/orderService4/v2
            http_methods:
                - ANY
            authentication:
                status: on
            # Rate limit use the key store for fast read write
            rate_limit:
                status: off
            # Circuit Breaker use the key store for fast read write
            circuit_breaker:
                status: off

        # Customers Internal Service
        - name: customer_service
          active: true
          proxy:
            listen_path: "/customer/v2/*"
            upstreams:
                balancing: random
                targets:
                    - target: https://httpbin.org/anything/customerService1/v2
                    - target: https://httpbin.org/anything/customerService2/v2
                    - target: https://httpbin.org/anything/customerService3/v2
                    - target: https://httpbin.org/anything/customerService4/v2
            http_methods:
                - GET
                - POST
                - PUT
                - DELETE
            authentication:
                status: on
            # Rate limit use the key store for fast read write
            rate_limit:
                status: off
            # Circuit Breaker use the key store for fast read write
            circuit_breaker:
                status: off

    # Log configs
    log:
        # Log level, it can be debug, info, warn, error, panic, fatal
        level: ${HELMET_LOG_LEVEL:-info}
        # Output can be stdout or abs path to log file /var/logs/helmet.log
        output: ${HELMET_LOG_OUTPUT:-stdout}
        # Format can be json
        format: ${HELMET_LOG_FORMAT:-json}
```

The run the `helmet` with `systemd`

```zsh
$ helmet server -c /path/to/config.yml
```


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Helmet is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/spacewalkio/helmet/releases) for changelogs for each release version of Helmet. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/spacewalkio/helmet/issues


## Security Issues

If you discover a security vulnerability within Helmet, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2021, SpaceWalk. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Helmet** is authored and maintained by [@SpaceWalk](http://github.com/spacewalkio).
