#!/bin/bash

function drifter {
    echo "Upgrade drifter ..."

    cd /etc/drifter
    mv config.prod.yml config.back.yml

    DRIFTER_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Spacemanio/Drifter/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/Spacemanio/Drifter/releases/download/v{$DRIFTER_LATEST_VERSION}/drifter_{$DRIFTER_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz

    rm config.prod.yml
    mv config.back.yml config.prod.yml

    systemctl restart drifter

    echo "drifter upgrade done!"
}

drifter
