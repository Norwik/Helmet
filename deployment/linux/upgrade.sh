#!/bin/bash

function helmet {
    echo "Upgrade helmet ..."

    cd /etc/helmet
    mv config.prod.yml config.back.yml

    DRIFTER_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Spacemanio/Helmet/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/Spacemanio/Helmet/releases/download/v{$DRIFTER_LATEST_VERSION}/helmet_{$DRIFTER_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz

    rm config.prod.yml
    mv config.back.yml config.prod.yml

    systemctl restart helmet

    echo "helmet upgrade done!"
}

helmet
