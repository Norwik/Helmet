#!/bin/bash

function helmet {
    echo "Upgrade Helmet ..."

    cd /etc/helmet
    mv config.prod.yml config.back.yml

    HELMET_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/spacewalkio/Helmet/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/spacewalkio/Helmet/releases/download/v{$HELMET_LATEST_VERSION}/helmet_{$HELMET_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz

    rm config.prod.yml
    mv config.back.yml config.prod.yml

    systemctl restart helmet

    echo "Helmet Upgrade Done!"
}

helmet
