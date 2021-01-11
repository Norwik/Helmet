<p align="center">
    <img src="/assets/logo.png?v=0.1.0" width="240" />
    <h3 align="center">Walnut</h3>
    <p align="center">A Lightweight Cloud Native API Gateway.</p>
    <p align="center">
        <a href="https://github.com/Clivern/Walnut/actions/workflows/build.yml">
            <img src="https://github.com/Clivern/Walnut/actions/workflows/build.yml/badge.svg">
        </a>
        <a href="https://github.com/Clivern/Walnut/actions">
            <img src="https://github.com/Clivern/Walnut/workflows/Release/badge.svg">
        </a>
        <a href="https://github.com/Clivern/Walnut/releases">
            <img src="https://img.shields.io/badge/Version-0.1.0-red.svg">
        </a>
        <a href="https://goreportcard.com/report/github.com/Clivern/Walnut">
            <img src="https://goreportcard.com/badge/github.com/Clivern/Walnut?v=0.1.0">
        </a>
        <a href="https://godoc.org/github.com/clivern/walnut">
            <img src="https://godoc.org/github.com/clivern/walnut?status.svg">
        </a>
        <a href="https://github.com/Clivern/Walnut/blob/master/LICENSE">
            <img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg">
        </a>
    </p>
</p>
<br/>


## Documentation

#### Linux Deployment Explained

Download [the latest walnut binary](https://github.com/Clivern/Walnut/releases). Make it executable from everywhere.

```zsh
$ export WALNUT_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Clivern/Walnut/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

$ curl -sL https://github.com/Clivern/Walnut/releases/download/v{$WALNUT_LATEST_VERSION}/walnut_{$WALNUT_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz
```

Then install `etcd` cluster or a single node! please [refer to etcd docs](https://etcd.io/docs/v3.5/) or bin directory inside this repository.

Create the configs file `config.yml` from `config.dist.yml`. Something like the following:

```yaml
# App configs
app:
    # Env mode (dev or prod)
    mode: ${WALNUT_APP_MODE:-dev}
    # HTTP port
    port: ${WALNUT_API_PORT:-8000}
    # Hostname
    hostname: ${WALNUT_API_HOSTNAME:-127.0.0.1}
    # TLS configs
    tls:
        status: ${WALNUT_API_TLS_STATUS:-off}
        pemPath: ${WALNUT_API_TLS_PEMPATH:-cert/server.pem}
        keyPath: ${WALNUT_API_TLS_KEYPATH:-cert/server.key}

    # API Configs
    api:
        key: ${WALNUT_API_KEY:-6c68b836-6f8e-465e-b59f-89c1db53afca}

    # Async Workers
    workers:
        # Queue max capacity
        buffer: ${WALNUT_WORKERS_CHAN_CAPACITY:-5000}
        # Number of concurrent workers
        count: ${WALNUT_WORKERS_COUNT:-4}

    # Runtime, Requests/Response and Walnut Metrics
    metrics:
        prometheus:
            # Route for the metrics endpoint
            endpoint: ${WALNUT_METRICS_PROM_ENDPOINT:-/metrics}

    # Application Database
    database:
        # Database driver
        driver: ${WALNUT_DB_DRIVER:-etcd}

        # Etcd Configs
        etcd:
            # Etcd database name or prefix
            databaseName: ${WALNUT_DB_ETCD_DB:-walnut}
            # Etcd username
            username: ${WALNUT_DB_ETCD_USERNAME:- }
            # Etcd password
            password: ${WALNUT_DB_ETCD_PASSWORD:- }
            # Etcd endpoints
            endpoints: ${WALNUT_DB_ETCD_ENDPOINTS:-http://127.0.0.1:2379}
            # Timeout in seconds
            timeout: 30

    # Log configs
    log:
        # Log level, it can be debug, info, warn, error, panic, fatal
        level: ${WALNUT_LOG_LEVEL:-info}
        # Output can be stdout or abs path to log file /var/logs/walnut.log
        output: ${WALNUT_LOG_OUTPUT:-stdout}
        # Format can be json
        format: ${WALNUT_LOG_FORMAT:-json}
```

The run the `walnut` with `systemd`

```zsh
$ walnut server -c /path/to/config.yml
```

#### To run the Admin Dashboard (Development Only):

Clone the project or your own fork:

```zsh
$ git clone https://github.com/Clivern/Walnut.git
```

Create the dashboard config file `web/.env` from `web/.env.dist`. Something like the following:

```
VUE_APP_API_URL=http://localhost:8080
```

Then you can either build or run the dashboard

```zsh
# Install npm packages
$ cd web
$ npm install
$ npm install -g npx

# Add api server url to frontend
$ echo "VUE_APP_API_URL=http://127.0.0.1:8000" > .env

$ cd ..

# Validate js code format
$ make check_ui_format

# Format UI
$ make format_ui

# Run Vuejs app
$ make serve_ui

# Build Vuejs app
$ make build_ui

# Any changes to the dashboard, must be reflected to cmd/pkged.go
# You can use these commands to do so
$ go get github.com/markbates/pkger/cmd/pkger
$ make package
```


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Walnut is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/walnut/releases) for changelogs for each release version of Walnut. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/walnut/issues


## Security Issues

If you discover a security vulnerability within Walnut, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2021, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**Walnut** is authored and maintained by [@clivern](http://github.com/clivern).
