#!/bin/bash

function helmet {
    echo "Installing helmet ..."

    apt-get install jq -y

    cd /etc/helmet

    HELMET_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Spacemanio/Helmet/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/Spacemanio/Helmet/releases/download/v{$HELMET_LATEST_VERSION}/helmet_{$HELMET_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz


    echo "[Unit]
Description=Helmet
Documentation=https://github.com/spacemanio/helmet

[Service]
ExecStart=/etc/helmet/helmet server -c /etc/helmet/config.prod.yml
Restart=on-failure
RestartSec=2

[Install]
WantedBy=multi-user.target" > /etc/systemd/system/helmet.service

    systemctl daemon-reload
    systemctl enable helmet.service
    systemctl start helmet.service

    echo "helmet installation done!"
}

helmet
