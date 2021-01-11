#!/bin/bash

function walnut {
    echo "Upgrade walnut ..."

    cd /etc/walnut
    mv config.prod.yml config.back.yml

    WALNUT_LATEST_VERSION=$(curl --silent "https://api.github.com/repos/Clivern/Walnut/releases/latest" | jq '.tag_name' | sed -E 's/.*"([^"]+)".*/\1/' | tr -d v)

    curl -sL https://github.com/Clivern/Walnut/releases/download/v{$WALNUT_LATEST_VERSION}/walnut_{$WALNUT_LATEST_VERSION}_Linux_x86_64.tar.gz | tar xz

    rm config.prod.yml
    mv config.back.yml config.prod.yml

    systemctl restart walnut

    echo "walnut upgrade done!"
}

walnut
