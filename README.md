<p align="center">
    <img src="/assets/logo1.png?v=0.1.0" width="240" />
    <h3 align="center">Drifter</h3>
    <p align="center">A Lightweight Cloud Native API Gateway.</p>
    <p align="center">
        <a href="https://github.com/Spacemanio/Drifter/actions/workflows/build.yml">
            <img src="https://github.com/Spacemanio/Drifter/actions/workflows/build.yml/badge.svg">
        </a>
        <a href="https://github.com/Spacemanio/Drifter/actions">
            <img src="https://github.com/Spacemanio/Drifter/workflows/Release/badge.svg">
        </a>
        <a href="https://github.com/Spacemanio/Drifter/releases">
            <img src="https://img.shields.io/badge/Version-0.1.0-red.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/Spacemanio/Drifter">
            <img src="https://goreportcard.com/badge/github.com/Spacemanio/Drifter?v=0.1.0">
        </a>
        <a href="https://godoc.org/github.com/spacemanio/drifter">
            <img src="https://godoc.org/github.com/spacemanio/drifter?status.svg">
        </a>
        <a href="https://github.com/Spacemanio/Drifter/blob/master/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg">
        </a>
    </p>
</p>
<br/>
<p align="center">
    <img src="/assets/chart.png?v=0.1.0" width="80%" />
</p>

Drifter is Cloud Native API Gateway that control who accesses your API whether from customer or other internal services. It also collect metrics about service calls count, latency, success rate and much more. here is some of the key features:

- Manage Service to Service Authentication, Authorization and Communication.
- Manage End User to Service Authentication, Authorization and Communication.
- Basic, API Key Based and OAuth2 Authentication Support.
- Multiple Backends Support with Load Balancing, Health Checks.
- OpenTracing support for Distributed tracing.
- Runtime Metrics exposed for Prometheus.
- CORS Support.
- HTTP/2 support.
- Rate Limiting Support.
- Caching Layer to make it even more faster.
- Lightweight, Easy to Deploy and Operate.


## Documentation

#### Linux Deployment

Download [the latest drifter binary](https://github.com/Spacemanio/Drifter/releases). Make it executable from everywhere.

```zsh
$ export DRIFTER_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Spacemanio/Drifter/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

$ curl -sL https://github.com/Spacemanio/Drifter/releases/download/v{$DRIFTER_LATEST_VERSION}/drifter_{$DRIFTER_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz
```

Then install `etcd` cluster or a single node! please [refer to etcd docs](https://etcd.io/docs/v3.5/) or bin directory inside this repository.

Create the configs file `config.yml` from `config.dist.yml`. Something like the following:

```yaml
# App configs
app:
    # App name
    name: ${DRIFTER_APP_NAME:-drifter}
    # Env mode (dev or prod)
    mode: ${DRIFTER_APP_MODE:-dev}
    # HTTP port
    port: ${DRIFTER_API_PORT:-8000}
    # Hostname
    hostname: ${DRIFTER_API_HOSTNAME:-127.0.0.1}
    # TLS configs
    tls:
        status: ${DRIFTER_API_TLS_STATUS:-off}
        crt_path: ${DRIFTER_API_TLS_PEMPATH:-cert/server.crt}
        key_path: ${DRIFTER_API_TLS_KEYPATH:-cert/server.key}

    # Global timeout
    timeout: ${DRIFTER_API_TIMEOUT:-50}

    # API Configs
    api:
        key: ${DRIFTER_API_KEY:-6c68b836-6f8e-465e-b59f-89c1db53afca}

    # Runtime, Requests/Response and Drifter Metrics
    metrics:
        prometheus:
            # Route for the metrics endpoint
            endpoint: ${DRIFTER_METRICS_PROM_ENDPOINT:-/_metrics}

    # Application Database
    database:
        # Database driver (sqlite3, mysql)
        driver: ${DRIFTER_DATABASE_DRIVER:-sqlite3}
        # Database Host
        host: ${DRIFTER_DATABASE_MYSQL_HOST:-localhost}
        # Database Port
        port: ${DRIFTER_DATABASE_MYSQL_PORT:-3306}
        # Database Name
        name: ${DRIFTER_DATABASE_MYSQL_DATABASE:-drifter.db}
        # Database Username
        username: ${DRIFTER_DATABASE_MYSQL_USERNAME:-root}
        # Database Password
        password: ${DRIFTER_DATABASE_MYSQL_PASSWORD:-root}

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
                auth_methods:
                    - 1
            rate_limit:
                status: off
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
                auth_methods:
                    - 1
            rate_limit:
                status: off
            circuit_breaker:
                status: off

    # Log configs
    log:
        # Log level, it can be debug, info, warn, error, panic, fatal
        level: ${DRIFTER_LOG_LEVEL:-info}
        # Output can be stdout or abs path to log file /var/logs/drifter.log
        output: ${DRIFTER_LOG_OUTPUT:-stdout}
        # Format can be json
        format: ${DRIFTER_LOG_FORMAT:-json}
```

The run the `drifter` with `systemd`

```zsh
$ drifter server -c /path/to/config.yml
```


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Drifter is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/spacemanio/drifter/releases) for changelogs for each release version of Drifter. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/spacemanio/drifter/issues


## Security Issues

If you discover a security vulnerability within Drifter, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2021, Spacemanio. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Drifter** is authored and maintained by [@clivern](http://github.com/clivern).
