#!/bin/bash

function drifter {
    echo "Installing drifter ..."

    apt-get install jq -y

    cd /etc/drifter

    DRIFTER_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Spacemanio/Drifter/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/Spacemanio/Drifter/releases/download/v{$DRIFTER_LATEST_VERSION}/drifter_{$DRIFTER_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz


    echo "[Unit]
Description=Drifter
Documentation=https://github.com/spacemanio/drifter

[Service]
ExecStart=/etc/drifter/drifter server -c /etc/drifter/config.prod.yml
Restart=on-failure
RestartSec=2

[Install]
WantedBy=multi-user.target" > /etc/systemd/system/drifter.service

    systemctl daemon-reload
    systemctl start drifter.service

    echo "drifter installation done!"
}

drifter
