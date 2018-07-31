#!/bin/sh

# From https://github.com/Boostport/kubernetes-vault/blob/master/deployments/quick-start/setup.sh#L3
# Check script dependencies
for command in jq kubectl vault
do
    if ! hash $command 2>/dev/null; then
        echo $command is not available, please install $command
        exit 1
    fi
done
