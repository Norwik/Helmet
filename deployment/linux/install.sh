#!/bin/bash

function mysql {
    echo "Installing MySQL ..."

    apt-get update

    echo "MySQL Installation Done!"
}

function redis {
    echo "Installing Redis ..."

    apt-get install redis-server

    echo "Redis Installation Done!"
}

function helmet {
    echo "Installing Helmet ..."

    apt-get install jq -y

    cd /etc/helmet

    HELMET_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/clevenio/Helmet/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/clevenio/Helmet/releases/download/v{$HELMET_LATEST_VERSION}/helmet_{$HELMET_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz


    echo "[Unit]
Description=Helmet
Documentation=https://github.com/clevenio/helmet

[Service]
ExecStart=/etc/helmet/helmet server -c /etc/helmet/config.prod.yml
Restart=on-failure
RestartSec=2

[Install]
WantedBy=multi-user.target" > /etc/systemd/system/helmet.service

    systemctl daemon-reload
    systemctl enable helmet.service
    systemctl start helmet.service

    echo "Helmet Installation Done!"
}

mysql
redis
helmet
