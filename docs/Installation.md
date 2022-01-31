### Installation

#### Ubuntu

- Install a `MySQL` server and create a database for `Helmet`.

```zsh
$ sudo apt update
$ sudo apt install mysql-server
$ sudo mysql_secure_installation
```

- Install a `Redis` Server.

```zsh
$ sudo apt install redis-server
```

- Install and configure `Helmet`.

```zsh
$ apt-get install jq -y

$ mkdir -p /etc/helmet
$ cd /etc/helmet

$ export LATEST_VERSION=$(curl --silent "https://api.github.com/repos/norwik/Helmet/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

$ curl -sL https://github.com/norwik/Helmet/releases/download/v{$LATEST_VERSION}/helmet_{$LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz
```

- Edit the `Helmet` config file `/etc/helmet/config.prod.yml` to configure the database and redis server. The following part

```yaml
# Application Database
database:
    # Database driver (sqlite3, mysql)
    driver: ${HELMET_DATABASE_DRIVER:-mysql}
    # Database Host
    host: ${HELMET_DATABASE_MYSQL_HOST:-localhost}
    # Database Port
    port: ${HELMET_DATABASE_MYSQL_PORT:-3306}
    # Database Name
    name: ${HELMET_DATABASE_MYSQL_DATABASE:-helmet}
    # Database Username
    username: ${HELMET_DATABASE_MYSQL_USERNAME:-root}
    # Database Password
    password: ${HELMET_DATABASE_MYSQL_PASSWORD:-root}

# Key Store Configs
key_store:
    # Cache Driver
    driver: ${HELMET_KV_DRIVER:-redis}
    # Redis Driver Configs
    redis:
        # Redis Address
        address: ${HELMET_KV_REDIS_ADDR:-localhost:6379}
        # Redis Password
        password: ${HELMET_KV_REDIS_PASSWORD:-}
        # Redis Database
        database: ${HELMET_KV_REDIS_DB:-0}
```

- Create a `Helmet` systemd service file.

```zsh
$ echo "[Unit]
Description=Helmet
Documentation=https://github.com/norwik/helmet

[Service]
ExecStart=/etc/helmet/helmet server -c /etc/helmet/config.prod.yml
Restart=on-failure
RestartSec=2

[Install]
WantedBy=multi-user.target" > /etc/systemd/system/helmet.service

$ systemctl daemon-reload
$ systemctl enable helmet.service
$ systemctl start helmet.service
```


#### Docker Compose
