<p align="center">
    <img src="/assets/logo.png?v=0.1.0" width="240" />
    <h3 align="center">Drifter</h3>
    <p align="center">A Lightweight Cloud Native API Gateway.</p>
    <p align="center">
        <a href="https://github.com/Clivern/Drifter/actions/workflows/build.yml">
            <img src="https://github.com/Clivern/Drifter/actions/workflows/build.yml/badge.svg">
        </a>
        <a href="https://github.com/Clivern/Drifter/actions">
            <img src="https://github.com/Clivern/Drifter/workflows/Release/badge.svg">
        </a>
        <a href="https://github.com/Clivern/Drifter/releases">
            <img src="https://img.shields.io/badge/Version-0.1.0-red.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/Clivern/Drifter">
            <img src="https://goreportcard.com/badge/github.com/Clivern/Drifter?v=0.1.0">
        </a>
        <a href="https://godoc.org/github.com/clivern/drifter">
            <img src="https://godoc.org/github.com/clivern/drifter?status.svg">
        </a>
        <a href="https://github.com/Clivern/Drifter/blob/master/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg">
        </a>
    </p>
</p>
<br/>

Drifter is Cloud Native API Gateway that control who accesses your API whether from customer or other internal services. It also collect metrics about service calls count, latency, success rate and much more. here is some of the key features:

- Manage service to service Authentication and Authorization.
- Manage user to service Authentication and Authorization.
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

Download [the latest drifter binary](https://github.com/Clivern/Drifter/releases). Make it executable from everywhere.

```zsh
$ export DRIFTER_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Clivern/Drifter/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

$ curl -sL https://github.com/Clivern/Drifter/releases/download/v{$DRIFTER_LATEST_VERSION}/drifter_{$DRIFTER_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz
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
        pemPath: ${DRIFTER_API_TLS_PEMPATH:-cert/server.pem}
        keyPath: ${DRIFTER_API_TLS_KEYPATH:-cert/server.key}

    # API Configs
    api:
        key: ${DRIFTER_API_KEY:-6c68b836-6f8e-465e-b59f-89c1db53afca}

    # Async Workers
    workers:
        # Queue max capacity
        buffer: ${DRIFTER_WORKERS_CHAN_CAPACITY:-5000}
        # Number of concurrent workers
        count: ${DRIFTER_WORKERS_COUNT:-4}

    # Runtime, Requests/Response and Drifter Metrics
    metrics:
        prometheus:
            # Route for the metrics endpoint
            endpoint: ${DRIFTER_METRICS_PROM_ENDPOINT:-/metrics}

    # Components Configs
    component:
        # Tracing Component
        tracing:
            # Status on or off
            status: ${DRIFTER_TRACING_STATUS:-on}
            # Tracing driver, jaeger supported so far
            driver: ${DRIFTER_TRACING_DRIVER:-jaeger}
            # Tracing backend URL
            collectorEndpoint: ${DRIFTER_TRACING_ENDPOINT:-http://jaeger.local:14268/api/traces}
            # Batch Size
            queueSize: ${DRIFTER_TRACING_QUEUE_SIZE:-20}

        # Profiler Component
        profiler:
            # Profiler Status
            status: ${DRIFTER_PROFILER_STATUS:-on}
            # Profiler Driver
            driver: ${DRIFTER_PROFILER_DRIVER:-log}

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

See the [Releases section of our GitHub project](https://github.com/clivern/drifter/releases) for changelogs for each release version of Drifter. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/drifter/issues


## Security Issues

If you discover a security vulnerability within Drifter, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2021, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Drifter** is authored and maintained by [@clivern](http://github.com/clivern).
